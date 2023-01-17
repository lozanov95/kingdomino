package server

import (
	"errors"
	"fmt"
	"io"
	"log"
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
		}
	}()
	log.Println("Started a game loop")
	buf := make([]byte, 1024)
	for {
		for _, player := range gr.Players {
			fmt.Println(player)
			n, err := player.Conn.Read(buf[0:])
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Println("error:", err)
				continue
			}
			fmt.Println(buf[:n])
		}
		time.Sleep(500 * time.Millisecond)
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
