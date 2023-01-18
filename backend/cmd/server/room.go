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
	Players     []*game.Player
	PlayerLimit int
	mux         sync.RWMutex
	Game        *game.Game
}

// Returns a new game room instance.
func NewGameRoom() *GameRoom {
	gr := &GameRoom{
		Players:     []*game.Player{},
		PlayerLimit: 2,
		mux:         sync.RWMutex{},
	}

	go gr.gameLoop()
	return gr
}

func (gr *GameRoom) gameLoop() {
	defer func() {
		for _, player := range gr.Players {
			player.Conn.Close()
			player.Connected = false
		}
		log.Println("game room closed")
	}()
	log.Println("Started a game loop")
	// buf := make([]byte, 1024)
	for len(gr.Players) < gr.PlayerLimit {
		time.Sleep(500 * time.Millisecond)
	}

	log.Println(gr.Players[0].Name, gr.Players[1].Name)

	for gr.Players[0].Connected && gr.Players[1].Connected {
		dices := gr.Game.RollDice()
		for _, player := range gr.Players {
			player.GameState <- game.GameState{Board: player.Board, BonusCard: player.BonusCard, Message: fmt.Sprintf("Player %s's turn to pick dice", gr.Players[0].Name), Dices: &dices, ID: player.Id}
		}

		log.Println("waiting for input")

		gr.handleDiceChoice(&dices, gr.Players[0])
		gr.handleDiceChoice(&dices, gr.Players[1])
		gr.handleDiceChoice(&dices, gr.Players[1])
		gr.handleDiceChoice(&dices, gr.Players[0])

		dices = gr.Game.RollDice()
		for _, player := range gr.Players {
			player.GameState <- game.GameState{Dices: &dices}
		}
		gr.handleDiceChoice(&dices, gr.Players[1])
		gr.handleDiceChoice(&dices, gr.Players[0])
		gr.handleDiceChoice(&dices, gr.Players[0])
		gr.handleDiceChoice(&dices, gr.Players[1])

		time.Sleep(10 * time.Second)
	}
}

func (gr *GameRoom) handleDiceChoice(d *[4]game.Badge, p *game.Player) {
	for {
		for _, player := range gr.Players {
			player.GameState <- game.GameState{Board: player.Board, Message: fmt.Sprintf("Player %s's turn to pick dice", p.Name), BonusCard: player.BonusCard, Dices: d, PlayerTurn: p.Id, ID: player.Id}
		}

		msg, err := p.GetInput()
		if err != nil {
			log.Println(err)
			return
		}
		choice, err := strconv.Atoi(string(msg))

		if err != nil || len(d) < choice || d[choice].Name == game.EMPTY {
			p.SendMessage("Invalid choice!")
			log.Println("Invalid choice")
			continue
		}

		d[choice].Name = game.EMPTY
		d[choice].Nobles = 0

		for _, player := range gr.Players {
			player.GameState <- game.GameState{Board: player.Board, Message: fmt.Sprintf("Player %s's turn to pick dice", p.Name), BonusCard: player.BonusCard, Dices: d, PlayerTurn: p.Id, ID: player.Id}
		}
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
