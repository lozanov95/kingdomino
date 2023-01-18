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
		for _, player := range gr.Players {
			player.Conn.Close()
			player.Connected = false
		}
		closeChan <- gr.ID
		log.Println("game room closed")
	}()
	log.Println("Opened a room")
	for len(gr.Players) < gr.PlayerLimit {
		time.Sleep(500 * time.Millisecond)
	}

	var dice *[4]game.Badge
	for _, p := range gr.Players {
		p.SendGameState(dice, "Connected!")
	}
	log.Println(gr.Players[0].Name, gr.Players[1].Name)

	for gr.Players[0].Connected && gr.Players[1].Connected {
		dice = gr.Game.RollDice()
		log.Println("waiting for input")

		gr.handleDiceChoice(dice, gr.Players[0])
		gr.handleDiceChoice(dice, gr.Players[1])
		gr.handleDiceChoice(dice, gr.Players[1])
		gr.handleDiceChoice(dice, gr.Players[0])

		dice = gr.Game.RollDice()

		gr.handleDiceChoice(dice, gr.Players[1])
		gr.handleDiceChoice(dice, gr.Players[0])
		gr.handleDiceChoice(dice, gr.Players[0])
		gr.handleDiceChoice(dice, gr.Players[1])
	}
}

func (gr *GameRoom) handleDiceChoice(d *[4]game.Badge, p *game.Player) {
	for {
		for _, player := range gr.Players {
			if !player.Connected {
				return
			}
			player.SendDice(d, fmt.Sprintf("Player %s's turn to pick dice", p.Name))
		}

		msg, err := p.GetInput()
		if err != nil {
			p.Connected = false
			log.Println("handle dice", err)
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
