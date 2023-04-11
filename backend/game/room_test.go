package game

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"golang.org/x/net/websocket"
)

type MockConn struct {
	inputs       []ClientPayload
	inputCounter int
}

func NewMockConn(cp []ClientPayload) *MockConn {
	return &MockConn{inputs: cp}
}

func NewMockPlayer(cp []ClientPayload) *Player {
	player := NewPlayer(NewMockConn(cp))
	go player.GameStateLoop()
	return player
}

// type Connectionable interface {
func (mc *MockConn) Read(msg []byte) (int, error) {
	input := mc.inputs[mc.inputCounter]
	mc.inputCounter = mc.inputCounter + 1
	payload, err := json.Marshal(input)
	if err != nil {
		log.Panic(err)
	}
	n := copy(msg, payload)
	return n, nil
}

func (mc *MockConn) Write([]byte) (int, error) {
	return 0, nil
}
func (mc *MockConn) SetReadDeadline(time.Time) error {
	return nil
}
func (mc *MockConn) Close() error {
	return nil
}

func TestJoinRoom(t *testing.T) {
	gr := NewGameRoom(make(chan string))
	p1 := NewPlayer(&websocket.Conn{})
	p2 := NewPlayer(&websocket.Conn{})
	p3 := NewPlayer(&websocket.Conn{})

	if err := gr.Join(p1); err != nil {
		t.Error(err)
	}
	if err := gr.Join(p2); err != nil {
		t.Error(err)
	}

	if err := gr.Join(p3); err != ErrGameRoomFull {
		t.Error("incorrect error on full room")
	}
	if len(gr.Players) != gr.PlayerLimit {
		t.Errorf("expected %d players, got %d", gr.PlayerLimit, len(gr.Players))
	}
}

func TestHandleDiceSelection(t *testing.T) {
	p1 := NewMockPlayer([]ClientPayload{
		{PlayerPower: PlayerPower{Use: true, Confirmed: true}}, {SelectedDie: 0}, {SelectedDie: 1},
	})
	p2 := NewMockPlayer([]ClientPayload{
		{SelectedDie: 1}, {SelectedDie: 3}, {SelectedDie: 2},
	})

	p1.AddBonus(Dice{Name: DOUBLEDOT, Nobles: 0})
	p1.AddBonus(Dice{Name: DOUBLEDOT, Nobles: 0})
	p1.AddBonus(Dice{Name: DOUBLEDOT, Nobles: 0})
	p1.AddBonus(Dice{Name: DOUBLEDOT, Nobles: 0})

	log.Println(p1.IsBonusCompleted(PWRPickTwoDice))
	gr := GameRoom{Players: []*Player{p1, p2}}
	gr.handleDicesSelection(&[]Dice{
		{Name: CHECKED},
		{Name: CHECKED},
		{Name: DOT},
		{Name: DOT}},
		p1, p2)

	if p1.IsBonusEligible(Dice{Name: DOUBLEDOT}) {
		t.Error("Expected the bonus to be ineligible")
	}
	if p1.Dices[0].Name != CHECKED && p1.Dices[1].Name != CHECKED {
		t.Errorf("Expected the selected dices to be %s and %s, got %s and %s instead\n",
			CHECKED, CHECKED, p1.Dices[0].Name, p1.Dices[1].Name)
	}
}
