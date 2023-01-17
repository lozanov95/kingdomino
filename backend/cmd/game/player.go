package game

import (
	"encoding/json"
	"log"

	"golang.org/x/net/websocket"
)

type Player struct {
	Name      string `json:"name"`
	Conn      *websocket.Conn
	board     *Board
	bonuscard *map[BadgeName]Bonus
}

// Creates a new player instance and returns a pointer to it.
func NewPlayer(jsonName []byte, conn *websocket.Conn) *Player {
	player := &Player{
		board:     NewBoard(),
		bonuscard: NewBonusMap(),
		Conn:      conn,
	}

	json.Unmarshal(jsonName, player)

	return player
}

// Increases the bonus of a specific card
func (p *Player) IncreaseBonus(b Badge) {
	if b.nobles != 0 {
		return
	}

	tmp := (*p.bonuscard)[b.name]
	tmp.Increment()
	(*p.bonuscard)[b.name] = tmp
}

func (p *Player) GetBoard() []byte {
	board, err := p.board.Json()
	if err != nil {
		log.Fatal(err)
	}
	return board
}
