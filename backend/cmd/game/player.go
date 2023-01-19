package game

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

var (
	ErrPlayerDisconnected = errors.New("the player has been disconnected")
)

const (
	TIMEOUT = 60 * time.Second
)

type DiePos struct {
	Cell int `json:"cell"`
	Row  int `json:"row"`
}

type BoardPlacementInput struct {
	Validate     bool
	PrevPosition DiePos
}

type GameState struct {
	ID            int64     `json:"id"`
	Message       string    `json:"message"`
	Board         *Board    `json:"board"`
	BonusCard     *BonusMap `json:"bonusCard"`
	Dices         *[4]Badge `json:"dices"`
	PlayerTurn    int64     `json:"playerTurn"`
	SelectedDices []Badge   `json:"selectedDice"`
}

type Player struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Conn      *websocket.Conn
	Board     *Board
	BonusCard *BonusMap
	Connected bool
	GameState chan GameState
	ClientMsg chan string
	Dices     []Badge
	mut       sync.RWMutex
}

// Creates a new player instance and returns a pointer to it.
func NewPlayer(jsonName []byte, conn *websocket.Conn) *Player {

	player := &Player{
		Id:        time.Now().UnixNano(),
		Name:      "",
		Conn:      conn,
		Board:     NewBoard(),
		BonusCard: NewBonusMap(),
		Connected: true,
		GameState: make(chan GameState),
		ClientMsg: make(chan string),
		Dices:     []Badge{},
		mut:       sync.RWMutex{},
	}

	json.Unmarshal(jsonName, player)

	return player
}

// Increases the bonus of a specific card
func (p *Player) IncreaseBonus(b Badge) {
	if b.Nobles != 0 {
		return
	}

	tmp := (*p.BonusCard)[b.Name]
	tmp.Increment()
	(*p.BonusCard)[b.Name] = tmp
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

func (p *Player) GameStateLoop() {
	for send := range p.GameState {
		msg, err := json.Marshal(send)
		if err != nil {
			if err == io.EOF {
				p.Connected = false
				return
			}
			log.Println(err)
		}
		p.Conn.Write(msg)
	}
}

func (p *Player) GetInput() ([]byte, error) {
	buf := make([]byte, 1024)

	err := p.Conn.SetReadDeadline(time.Now().Add(TIMEOUT))
	if err != nil {
		return nil, err
	}
	n, err := p.Conn.Read(buf[0:])
	if err != nil {
		p.Connected = false
		panic(ErrPlayerDisconnected)

	}

	return buf[:n], nil
}

func (p *Player) SendMessage(message string) {
	p.GameState <- GameState{Message: message, SelectedDices: p.Dices}
}

func (p *Player) SendDice(d *[4]Badge, m string) {
	p.GameState <- GameState{Dices: d, Message: m, SelectedDices: p.Dices}
}

func (p *Player) SendGameState(d *[4]Badge, m string) {
	p.GameState <- GameState{
		ID:            p.Id,
		Message:       m,
		Board:         p.Board,
		BonusCard:     p.BonusCard,
		Dices:         d,
		SelectedDices: p.Dices,
		// PlayerTurn: 0,
	}
}

func (p *Player) AddDice(d Badge) {
	p.mut.Lock()
	defer p.mut.Unlock()
	p.Dices = append(p.Dices, d)
}

func (p *Player) ClearDice() {
	p.mut.Lock()
	defer p.mut.Unlock()
	p.Dices = make([]Badge, 0)
}

func (p *Player) PlaceDomino() {
	var d [4]Badge

	p.SendMessage("Select the dice that you want to place")
	choice := p.getSelectedDominoChoice()
	prevPos := p.placeOnBoard(choice, BoardPlacementInput{})
	p.SendGameState(&d, "")

	p.SendMessage("Select the dice that you want to place")
	choice = p.getSelectedDominoChoice()
	p.placeOnBoard(choice, BoardPlacementInput{Validate: true, PrevPosition: prevPos})
	p.SendGameState(&d, "")
}

func (p *Player) getSelectedDominoChoice() int {
	for {
		var choice int
		msg, err := p.GetInput()
		if err != nil {
			return choice
		}

		choice, err = strconv.Atoi(string(msg))

		if err != nil || len(p.Dices) < choice || p.Dices[choice].Name == EMPTY {
			p.SendMessage("Invalid choice!")
			log.Println("Invalid choice")
			continue
		}

		return choice
	}
}

func (p *Player) placeOnBoard(choice int, b BoardPlacementInput) DiePos {
	pos, err := p.getBoardPlacementInput(b)
	if err != nil {
		return DiePos{}
	}

	p.Board[pos.Row][pos.Cell] = p.Dices[choice]
	newDices := p.Dices[:choice]
	newDices = append(newDices, p.Dices[choice+1:]...)
	p.Dices = newDices
	return pos
}

func (p *Player) getBoardPlacementInput(bpi BoardPlacementInput) (DiePos, error) {
	p.SendMessage("Select the place on the board that you want to place it on")
	for p.Connected {
		msg, err := p.GetInput()
		if err != nil {
			if err == io.EOF {
				p.Connected = false
			}
			log.Println(err)
			return DiePos{}, err
		}
		var pos DiePos
		err = json.Unmarshal(msg, &pos)
		if err != nil {
			if err == io.EOF {
				p.Connected = false
			}
			log.Println(err)
			return DiePos{}, err
		}

		if !p.Board.IsValidPlacementPos(pos.Row, pos.Cell) || (bpi.Validate && !bpi.IsValid(&pos)) {
			p.SendMessage("Invalid position. Please select a new position")
			continue
		}

		return pos, nil
	}
	return DiePos{}, io.EOF
}

func (bpi *BoardPlacementInput) IsValid(newPos *DiePos) bool {
	if (newPos.Row-1 == bpi.PrevPosition.Row && newPos.Cell == bpi.PrevPosition.Cell) ||
		(newPos.Row+1 == bpi.PrevPosition.Row && newPos.Cell == bpi.PrevPosition.Cell) ||
		(newPos.Row == bpi.PrevPosition.Row && newPos.Cell-1 == bpi.PrevPosition.Cell) ||
		(newPos.Row == bpi.PrevPosition.Row && newPos.Cell+1 == bpi.PrevPosition.Cell) {
		return true
	}

	return false
}
