package game

import (
	"encoding/json"
	"log"
	"strings"

	"golang.org/x/net/websocket"
)

type Player struct {
	Name  string `json:"name"`
	Conn  *websocket.Conn
	board *Board
	// bonuscard *map[BadgeName]Bonus
	bonuscard *BonusMap
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

func (p *Player) GetBonusCard() []byte {
	card, err := p.bonuscard.MarshalJSON()
	if err != nil {
		log.Println(err)
		return nil
	}

	return card
}

func (p *Player) GetState() []byte {
	var sb strings.Builder
	sb.WriteString("{")
	board := strings.Replace(string(p.GetBoard()), "{", "", 1)
	board = strings.TrimSuffix(board, "}")
	bonusCard := strings.Replace(string(p.GetBonusCard()), "{", "", 1)
	bonusCard = strings.TrimSuffix(bonusCard, "}")
	sb.WriteString(board)
	sb.WriteString(",")
	sb.WriteString(bonusCard)
	sb.WriteString("}")

	return []byte(sb.String())
}
