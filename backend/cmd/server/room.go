package server

import (
	"errors"
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
			player.Connected = false
		}
	}()
	log.Println("Started a game loop")
	// buf := make([]byte, 1024)
	for {
		if len(gr.Players) < gr.PlayerLimit {
			continue
		}
		log.Println(gr.Players[0].Name, gr.Players[1].Name)
		for _, player := range gr.Players {
			player.GameState <- game.GameState{Board: player.Board, BonusCard: player.BonusCard, Message: "yo"}
			// n, err := player.Conn.Read(buf[0:])
			// if err != nil {
			// 	if err == io.EOF {
			// 		player.Connected = false
			// 		log.Printf("player %s disconnected\n", player.Name)
			// 		return
			// 	}
			// 	log.Println("error:", err)
			// 	continue
			// }
			// fmt.Println(buf[:n])
			// player.ClientMsg <- string(buf[:n])
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
