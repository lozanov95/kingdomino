package game

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

type GameState struct {
	Message   string    `json:"message"`
	Board     *Board    `json:"board"`
	BonusCard *BonusMap `json:"bonusCard"`
	Dices     *[4]Badge `json:"dices"`
}

type Player struct {
	Name      string `json:"name"`
	Conn      *websocket.Conn
	Board     *Board
	BonusCard *BonusMap
	Connected bool
	GameState chan GameState
	ClientMsg chan string
}

// Creates a new player instance and returns a pointer to it.
func NewPlayer(jsonName []byte, conn *websocket.Conn) *Player {
	player := &Player{
		Board:     NewBoard(),
		BonusCard: NewBonusMap(),
		Conn:      conn,
		Connected: true,
		GameState: make(chan GameState),
		ClientMsg: make(chan string),
	}

	json.Unmarshal(jsonName, player)

	return player
}

// Increases the bonus of a specific card
func (p *Player) IncreaseBonus(b Badge) {
	if b.nobles != 0 {
		return
	}

	tmp := (*p.BonusCard)[b.name]
	tmp.Increment()
	(*p.BonusCard)[b.name] = tmp
}

func (p *Player) GetBoard() []byte {
	board, err := p.Board.Json()
	if err != nil {
		log.Fatal(err)
	}
	return board
}

func (p *Player) GetBonusCard() []byte {
	card, err := p.BonusCard.MarshalJSON()
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

func (p *Player) SendGameState() {
	for p.Connected {
		select {
		case send := <-p.GameState:
			msg, err := json.Marshal(send)
			if err != nil {
				if err == io.EOF {
					p.Connected = false
					return
				}
				log.Println(err)
			}
			p.Conn.Write(msg)
			// fmt.Println(p.Name, "sending", string(msg))
		case receive := <-p.ClientMsg:
			log.Println(receive)
		}
		fmt.Println(p.Name, p.Conn.RemoteAddr())
		time.Sleep(1000 * time.Millisecond)
	}
}
