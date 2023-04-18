package game

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

// Represents the payload that the client sends to the server
type ClientPayload struct {
	// Name of the player
	Name string `json:"name"`

	// The placement position of the die
	DiePos DiePos `json:"boardPosition"`

	// The die that the user have selected
	SelectedDie int `json:"selectedDie"`

	// Response containing the user's choice to use/ not use the power
	PlayerPower `json:"playerPower"`
}

// Represents the Game State that is being send to the player
type GameState struct {
	// The room ID
	ID int64 `json:"id"`

	// Message to the player
	Message string `json:"message"`

	// Player's domino board
	Board *Board `json:"board"`

	// The bonus card that keeps track of wizard powers
	BonusCard *BonusMap `json:"bonusCard"`

	Dices *[]DiceResult `json:"dices"`

	// The available wizard power to the player
	PlayerPower `json:"playerPower"`

	Scoreboards []Scoreboard `json:"scoreboards"`
}

type GameRoom struct {
	ID          string
	Players     []*Player
	PlayerLimit int
	mux         sync.RWMutex
	Game        *Game
}

type DiceResult struct {
	*Dice      `json:"dice"`
	IsSelected bool  `json:"isSelected"`
	PlayerId   int64 `json:"playerId"`
	IsPlaced   bool  `json:"isPlaced"`
}

var (
	ErrGameRoomFull = errors.New("the game room is full")
)

const (
	// The duration after which the player will be kicked for inactivity.
	TIMEOUT = 60 * time.Minute
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
	p1_sb := gr.Players[0].CalculateScore()
	p2_sb := gr.Players[1].CalculateScore()

	if p1_sb.TotalScore > p2_sb.TotalScore {
		gr.Players[0].SendScoreboard(&p1_sb, &p2_sb, "Game OVER! You WON!")
		gr.Players[1].SendScoreboard(&p2_sb, &p1_sb, "Game OVER! You LOST!")
	} else if p1_sb.TotalScore < p2_sb.TotalScore {
		gr.Players[0].SendScoreboard(&p1_sb, &p2_sb, "Game OVER! You WON!")
		gr.Players[1].SendScoreboard(&p2_sb, &p1_sb, "Game OVER! You LOST!")
	} else {
		gr.Players[0].SendScoreboard(&p1_sb, &p2_sb, "Game OVER! The game ended in a DRAW!")
		gr.Players[1].SendScoreboard(&p2_sb, &p1_sb, "Game OVER! The game ended in a DRAW!")
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

	for _, p := range gr.Players {
		p.SendGameState(nil, "Connected!")
	}
	log.Println("started room with", gr.Players[0].Name, "and", gr.Players[1].Name)

	gr.gameLoop()
	gr.score()
	time.Sleep(10 * time.Second)
}

// Checks if the requirements for a valid running game are satisfied
func (gr *GameRoom) shouldLoopContinue() bool {
	return gr.Players[0].Connected &&
		gr.Players[1].Connected &&
		(gr.Players[0].IsAnyPlacementPossible() ||
			gr.Players[1].IsAnyPlacementPossible())
}

// Handles the main game loop - selecting dice and placing dominos
func (gr *GameRoom) gameLoop() {
	var wg sync.WaitGroup
	for {
		if !gr.shouldLoopContinue() {
			break
		}
		gr.playGameRound(&wg, gr.Players[0], gr.Players[1])

		if !gr.shouldLoopContinue() {
			break
		}
		gr.playGameRound(&wg, gr.Players[1], gr.Players[2])
	}
}

// Handles a full round - roll dice, select dice, place domino
func (gr *GameRoom) playGameRound(wg *sync.WaitGroup, p1, p2 *Player) {
	dice := gr.Game.RollDice()
	gr.handleDicesSelection(dice, gr.Players[0], gr.Players[1])
	wg.Add(len(gr.Players))
	for _, player := range gr.Players {
		go gr.handlePlaceDomino(wg, player, dice)
	}
	wg.Wait()
}

// Handles the situation where two players take turns to choose a die
func (gr *GameRoom) handleDicesSelection(dice *[]DiceResult, p1, p2 *Player) {
	p1.handleUseAddNoblePower()
	p2.handleUseAddNoblePower()
	p1.handleAddDomainPointsPower()
	p2.handleAddDomainPointsPower()

	if p1.IsBonusUsable(PWRPickTwoDice) {
		p1.SendPlayerPowerPrompt(dice, PlayerPower{Type: PWRPickTwoDice, Description: "Pick two dices immediately."})
		p2.SendGameState(dice, "Waiting for your opponent to decide if they want to use a wizard power")
		payload := p1.GetInput()

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

// Handles the situation of a single player choosing a die
func (gr *GameRoom) handleDiceChoice(d *[]DiceResult, p, p2 *Player) {
	for {
		for _, player := range gr.Players {
			if !player.Connected {
				return
			}
			player.SendGameState(d, fmt.Sprintf("Player %s's turn to pick dice", p.GetName()))
		}

		choice := -1
		if p.IsBonusUsable(PWRSelectDieSideOfChoice) {
			p.SendPlayerPowerPrompt(d,
				PlayerPower{
					Type:        PWRSelectDieSideOfChoice,
					Description: "You can turn one of the dice in your domino around so that it shows any face of your choice"},
			)

			if p.GetPlayerPowerChoice() {
				p.UsePower(PWRSelectDieSideOfChoice)
				p.SendMessage("Select which die you want to turn")
				choice = func(d *[]DiceResult) int {
					for {
						payload := p.GetInput()
						if (*d)[payload.SelectedDie].Name == EMPTY {
							p.SendMessage("Invalid selection! Please choose another die")
							continue
						}

						return payload.SelectedDie
					}
				}(d)

				dice := gr.Game.GetDieAllSides(choice)
				p.SendDice(&dice, "Choose die")

				payload := p.GetInput()
				selectedDie := gr.Game.dices[choice][payload.SelectedDie]
				(*d)[choice] = *NewDiceResult(selectedDie)

				if selectedDie.Name == QUESTIONMARK {
					handleQuestionmark(d, choice, p)
					return
				}

				p.SelectDie(&(*d)[choice])
				if p.IsBonusCompleted(getBonusType(selectedDie.Name)) {
					p2.SetBonusIneligible(selectedDie)
				}

				for _, player := range gr.Players {
					if !player.Connected {
						return
					}
					player.SendGameState(d, fmt.Sprintf("Player %s's turn to pick dice", p.GetName()))
				}

				return
			}
		}

		if choice == -1 {
			payload := p.GetInput()
			choice = payload.SelectedDie
		}

		if !isDicePickChoiceValid(d, choice) {
			p.SendMessage("Invalid choice!")
			log.Println("Invalid choice")
			continue
		}

		if (*d)[choice].Name == QUESTIONMARK {
			handleQuestionmark(d, choice, p)
			return
		}

		p.SelectDie(&(*d)[choice])
		if p.IsBonusCompleted(getBonusType((*d)[choice].Dice.Name)) {
			p2.SetBonusIneligible(*(*d)[choice].Dice)
		}

		return
	}
}

// Handles the place domino section
func (gr *GameRoom) handlePlaceDomino(wg *sync.WaitGroup, player *Player, dice *[]DiceResult) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			player.Connected = false
		}
	}()
	defer wg.Done()
	if !player.IsBonusUsable(PWRSeparateDominos) && player.IsValidPlacementPossible() {
		player.PlaceDomino(dice)
		return
	}

	if !player.IsBonusUsable(PWRSeparateDominos) || !player.IsThereAFreeSpot() {
		return
	}

	player.SendPlayerPowerPrompt(dice,
		PlayerPower{
			Type:        PWRSeparateDominos,
			Description: "You can separate your dice to fill in your map. Each die must respect the Connection Rules",
		})
	payload := player.GetInput()

	if payload.Use && player.IsThereAFreeSpot() {
		player.UsePower(PWRSeparateDominos)
		player.PlaceSeparatedDomino(dice)
		return
	}
	player.PlaceDomino(dice)
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

func isDicePickChoiceValid(availableDice *[]DiceResult, choice int) bool {
	if choice < 0 ||
		choice >= len(*availableDice) ||
		(*availableDice)[choice].IsSelected {

		return false
	}

	return true
}

func isDicePlaceChoiceValid(availableDice *[]DiceResult, choice int, playerId int64) bool {
	if choice < 0 ||
		choice >= len(*availableDice) ||
		(*availableDice)[choice].IsPlaced ||
		(*availableDice)[choice].PlayerId != playerId {

		return false
	}

	return true
}

func handleQuestionmark(d *[]DiceResult, initialChoice int, p *Player) {
	newDice := &[]DiceResult{
		*NewDiceResult(Dice{Name: DOT}),
		*NewDiceResult(Dice{Name: LINE}),
		*NewDiceResult(Dice{Name: DOUBLEDOT}),
		*NewDiceResult(Dice{Name: DOUBLELINE}),
		*NewDiceResult(Dice{Name: FILLED}),
		*NewDiceResult(Dice{Name: CHECKED}),
	}

	p.SendDice(newDice, "Please select the type of badge that you need")
	for {
		payload := p.GetInput()
		choice := payload.SelectedDie

		if !isDicePickChoiceValid(newDice, choice) {
			p.SendMessage("Invalid choice!")
			log.Println("Invalid choice")
			continue
		}
		(*d)[initialChoice] = (*newDice)[choice]
		p.SelectQuestionmarkDie(&(*d)[initialChoice])

		return
	}
}
