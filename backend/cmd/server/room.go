package server

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/lozanov95/kingdomino/backend/cmd/game"
)

var (
	ErrGameRoomFull = errors.New("the game room is full")
)

type GameRoom struct {
	ID          string
	Players     []*game.Player
	PlayerLimit int
	mux         sync.RWMutex
	Game        *game.Game
}

// Returns a new game room instance.
func NewGameRoom(closeChan chan string) *GameRoom {
	id := strconv.Itoa(int(time.Now().UnixMicro()))
	gr := &GameRoom{
		ID:          id,
		Players:     []*game.Player{},
		PlayerLimit: 2,
		mux:         sync.RWMutex{},
	}

	go gr.gameLoop(closeChan)
	return gr
}

func (gr *GameRoom) gameLoop(closeChan chan<- string) {
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

	var dice *[]game.Badge
	for _, p := range gr.Players {
		p.SendGameState(dice, "Connected!")
	}
	log.Println("started room with", gr.Players[0].Name, "and", gr.Players[1].Name)

	var wg sync.WaitGroup

	for gr.Players[0].Connected && gr.Players[1].Connected && (gr.Players[0].IsValidPlacementPossible() || gr.Players[1].IsValidPlacementPossible()) {

		dice = gr.Game.RollDice()

		gr.handleDicesRound(dice, gr.Players[0], gr.Players[1])

		for _, player := range gr.Players {
			wg.Add(1)
			go func(player *game.Player) {
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
		gr.handleDicesRound(dice, gr.Players[1], gr.Players[0])
		for _, player := range gr.Players {
			wg.Add(1)
			go func(player *game.Player) {
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
	gr.Players[0].SendMessage("Game OVER!")
	gr.Players[1].SendMessage("Game OVER!")
	time.Sleep(10 * time.Second)
}

func (gr *GameRoom) handleDicesRound(dice *[]game.Badge, p1, p2 *game.Player) {
	for _, player := range gr.Players {
		player.ClearDice()
	}

	gr.handleDiceChoice(dice, p1, p2)
	gr.handleDiceChoice(dice, p2, p1)
	gr.handleDiceChoice(dice, p2, p1)
	gr.handleDiceChoice(dice, p1, p2)
}

func (gr *GameRoom) handleDiceChoice(d *[]game.Badge, p *game.Player, p2 *game.Player) {
	for {
		for _, player := range gr.Players {
			if !player.Connected {
				return
			}
			player.SendGameState(d, fmt.Sprintf("Player %s's turn to pick dice", p.Name))
		}

		payload, err := p.GetInput()
		choice := payload.SelectedDie

		if err != nil || choice < 0 || len((*d)) < choice || (*d)[choice].Name == game.EMPTY {
			p.SendMessage("Invalid choice!")
			log.Println("Invalid choice")
			continue
		}

		if (*d)[choice].Name == game.QUESTIONMARK {
			newDice := &[]game.Badge{
				{Name: game.DOT},
				{Name: game.LINE},
				{Name: game.DOUBLEDOT},
				{Name: game.DOUBLELINE},
				{Name: game.FILLED},
				{Name: game.CHECKED},
			}
			p.SendDice(newDice, "Please select the type of badge that you need!")
			for {
				payload, err := p.GetInput()
				newChoice := payload.SelectedDie

				if err != nil || newChoice < 0 || len((*newDice)) < newChoice || (*newDice)[newChoice].Name == game.EMPTY {
					p.SendMessage("Invalid choice!")
					log.Println("Invalid choice")
					continue
				}

				p.AddDice((*newDice)[newChoice])

				(*d)[choice].Name = game.EMPTY
				(*d)[choice].Nobles = 0

				return
			}
		}

		p.AddDice((*d)[choice])
		p.BonusCard.AddBonus((*d)[choice])
		dName := (*d)[choice].Name
		if !(*p.BonusCard)[dName].Eligible {
			p2Bonus := (*p2.BonusCard)[dName]
			p2Bonus.Eligible = false
			(*p2.BonusCard)[dName] = p2Bonus
		}
		(*d)[choice].Name = game.EMPTY
		(*d)[choice].Nobles = 0

		return
	}
}

// Joins a game room
// Returns ErrGameRoomFull if there is no space in the room.
func (gr *GameRoom) Join(p *game.Player) error {
	if gr.IsFull() {
		return ErrGameRoomFull
	}

	gr.mux.Lock()
	defer gr.mux.Unlock()
	gr.Players = append(gr.Players, p)

	if gr.PlayerLimit > len(gr.Players) {
		return nil
	}

	gr.Game = game.NewGame(gr.Players[0], gr.Players[1])

	return nil
}

// Checks if a the game room is full
func (gr *GameRoom) IsFull() bool {
	return len(gr.Players) >= gr.PlayerLimit
}
