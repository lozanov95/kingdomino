package game

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

type ClientPayload struct {
	Name        string `json:"name"`
	DiePos      DiePos `json:"boardPosition"`
	SelectedDie int    `json:"selectedDie"`
	PlayerPower `json:"playerPower"`
}

type GameState struct {
	ID            int64     `json:"id"`
	Message       string    `json:"message"`
	Board         *Board    `json:"board"`
	BonusCard     *BonusMap `json:"bonusCard"`
	Dices         *[]Badge  `json:"dices"`
	GameTurn      GameTurn  `json:"gameTurn"`
	PlayerPower   `json:"playerPower"`
	SelectedDices []Badge `json:"selectedDice"`
}

type GameRoom struct {
	ID          string
	Players     []*Player
	PlayerLimit int
	mux         sync.RWMutex
	Game        *Game
}

type GameTurn int64

var (
	ErrGameRoomFull = errors.New("the game room is full")
)

const (
	// The duration after which the player will be kicked for inactivity.
	TIMEOUT = 2 * time.Minute
)

const (
	// The game is waiting for both players to connect
	GTWaitingForPlayers GameTurn = iota

	// The game is waiting for both players to pick dice
	GTPickDice

	// The game is waiting for both players to place domino
	GTPlaceDomino

	// The game is waiting for the Magic Powers selection
	GTUseMagicPowers

	// Waiting for player to conduct their turn
	GTWaitingPlayerTurn

	// The game is over
	GTGameOver
)

// Returns a new game room instance.
func NewGameRoom(closeChan chan string) *GameRoom {
	id := strconv.Itoa(int(time.Now().UnixMicro()))
	gr := &GameRoom{
		ID:          id,
		Players:     []*Player{},
		PlayerLimit: 2,
		mux:         sync.RWMutex{},
	}

	go gr.roomLoop(closeChan)
	return gr
}

// Calculates the score and sends message to the players
func (gr *GameRoom) score() {
	p1_score := gr.Players[0].CalculateScore()
	p2_score := gr.Players[1].CalculateScore()

	if p1_score > p2_score {
		gr.Players[0].SendMessage("Game OVER! You WON!")
		gr.Players[1].SendMessage("Game OVER! You LOST!")
	} else if p1_score < p2_score {
		gr.Players[1].SendMessage("Game OVER! You WON!")
		gr.Players[0].SendMessage("Game OVER! You LOST!")
	} else {
		gr.Players[0].SendMessage("Game OVER! The game ended in a DRAW!")
		gr.Players[1].SendMessage("Game OVER! The game ended in a DRAW!")
	}
}

func (gr *GameRoom) roomLoop(closeChan chan<- string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic in game loop", err)
		}
		for _, player := range gr.Players {
			player.Conn.Close()
			player.Connected = false
		}
		closeChan <- gr.ID
	}()

	log.Println("Opened a room")
	for len(gr.Players) < gr.PlayerLimit {
		time.Sleep(500 * time.Millisecond)
	}

	var dice *[]Badge
	for _, p := range gr.Players {
		p.SendGameState(dice, "Connected!", GTWaitingForPlayers)
	}
	log.Println("started room with", gr.Players[0].Name, "and", gr.Players[1].Name)

	gr.gameLoop(dice)
	gr.score()
	time.Sleep(10 * time.Second)
}

// Checks if the requirements for a valid running game are satisfied
func (gr *GameRoom) shouldLoopContinue() bool {
	return gr.Players[0].Connected &&
		gr.Players[1].Connected &&
		(gr.Players[0].IsValidPlacementPossible() ||
			gr.Players[1].IsValidPlacementPossible())
}

// Handles the main game loop - selecting dice and placing dominos
func (gr *GameRoom) gameLoop(dice *[]Badge) {
	var wg sync.WaitGroup
	for gr.shouldLoopContinue() {
		dice = gr.Game.RollDice()
		gr.handleDicesSelection(dice, gr.Players[0], gr.Players[1])

		for _, player := range gr.Players {
			wg.Add(1)
			go func(player *Player) {
				defer func() {
					if err := recover(); err != nil {
						log.Println(err)
						player.Connected = false
					}
				}()
				defer wg.Done()
				if player.IsValidPlacementPossible() {
					player.PlaceDomino(dice)
				}
			}(player)
		}

		wg.Wait()
		dice := gr.Game.RollDice()
		gr.handleDicesSelection(dice, gr.Players[1], gr.Players[0])
		for _, player := range gr.Players {
			wg.Add(1)
			go func(player *Player) {
				defer func() {
					if err := recover(); err != nil {
						log.Println(err)
						player.Connected = false
					}
				}()
				defer wg.Done()
				if player.IsValidPlacementPossible() {
					player.PlaceDomino(dice)
				}
			}(player)
		}

		wg.Wait()

		for _, player := range gr.Players {
			player.ClearDice()
		}
	}
}

func (gr *GameRoom) handleDicesSelection(dice *[]Badge, p1, p2 *Player) {
	for _, player := range gr.Players {
		player.ClearDice()
	}

	if p1.IsBonusCompleted(PWRPickTwoDice) {
		p1.SendPlayerPowerPrompt(PlayerPower{Type: PWRPickTwoDice, Description: "Pick two dices immediately."})
		p2.SendGameState(dice, "Waiting for your opponent to decide if they want to use a wizard power", GTWaitingPlayerTurn)
		payload, err := p1.GetInput()
		if err != nil {
			log.Println(err)
			p1.Disconnect()
			return
		}

		if payload.PlayerPower.Use {
			p1.UsePower(PWRPickTwoDice)
			gr.handleDiceChoice(dice, p1, p2)
			gr.handleDiceChoice(dice, p1, p2)
			gr.handleDiceChoice(dice, p2, p1)
			gr.handleDiceChoice(dice, p2, p1)
			return
		}
	}

	gr.handleDiceChoice(dice, p1, p2)
	gr.handleDiceChoice(dice, p2, p1)
	gr.handleDiceChoice(dice, p2, p1)
	gr.handleDiceChoice(dice, p1, p2)

}

func (gr *GameRoom) handleDiceChoice(d *[]Badge, p, p2 *Player) {
	for {
		for _, player := range gr.Players {
			if !player.Connected {
				return
			}
			player.SendGameState(d, fmt.Sprintf("Player %s's turn to pick dice", p.GetName()), GTPickDice)
		}

		payload, err := p.GetInput()
		choice := payload.SelectedDie

		if err != nil || choice < 0 || len((*d)) < choice || (*d)[choice].Name == EMPTY {
			p.SendMessage("Invalid choice!")
			log.Println("Invalid choice")
			continue
		}

		if (*d)[choice].Name == QUESTIONMARK {
			newDice := &[]Badge{
				{Name: DOT},
				{Name: LINE},
				{Name: DOUBLEDOT},
				{Name: DOUBLELINE},
				{Name: FILLED},
				{Name: CHECKED},
			}
			p.SendDice(newDice, "Please select the type of badge that you need")
			for {
				payload, err := p.GetInput()
				newChoice := payload.SelectedDie

				if err != nil || newChoice < 0 || len((*newDice)) < newChoice || (*newDice)[newChoice].Name == EMPTY {
					p.SendMessage("Invalid choice!")
					log.Println("Invalid choice")
					continue
				}

				p.AddDice((*newDice)[newChoice])

				(*d)[choice].Name = EMPTY
				(*d)[choice].Nobles = 0

				return
			}
		}

		selectedDie := (*d)[choice]
		p.AddDice(selectedDie)
		p.AddBonus(selectedDie)
		if p.IsBonusCompleted(getBonusType(selectedDie.Name)) {
			p2.SetBonusIneligible(selectedDie)
		}

		(*d)[choice].Name = EMPTY
		(*d)[choice].Nobles = 0

		return
	}
}

// Joins a game room
// Returns ErrGameRoomFull if there is no space in the room.
func (gr *GameRoom) Join(p *Player) error {
	if gr.IsFull() {
		return ErrGameRoomFull
	}

	gr.mux.Lock()
	defer gr.mux.Unlock()
	gr.Players = append(gr.Players, p)

	if gr.PlayerLimit > len(gr.Players) {
		return nil
	}

	gr.Game = NewGame(gr.Players[0], gr.Players[1])

	return nil
}

// Checks if a the game room is full
func (gr *GameRoom) IsFull() bool {
	return len(gr.Players) >= gr.PlayerLimit
}
