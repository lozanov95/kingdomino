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
	dices := &[]DiceResult{
		*NewDiceResult(&Dice{Name: CHECKED}),
		*NewDiceResult(&Dice{Name: CHECKED}),
		*NewDiceResult(&Dice{Name: DOT}),
		*NewDiceResult(&Dice{Name: DOT})}
	gr.handleDicesSelection(dices,
		p1, p2)

	if p1.IsBonusEligible(Dice{Name: DOUBLEDOT}) {
		t.Error("Expected the bonus to be ineligible")
	}
	if (!(*dices)[0].IsSelected || (*dices)[0].PlayerId != p1.Id) ||
		(!(*dices)[1].IsSelected || (*dices)[1].PlayerId != p1.Id) ||
		(!(*dices)[3].IsSelected || (*dices)[3].PlayerId != p2.Id) ||
		(!(*dices)[2].IsSelected || (*dices)[2].PlayerId != p2.Id) {
		t.Error("The dice selection was incorrect")
	}
}

func TestHandleQuestionmark(t *testing.T) {
	p1 := NewMockPlayer([]ClientPayload{{SelectedDie: 0}})
	dr := []DiceResult{
		{Dice: &Dice{Name: DOT}, PlayerId: p1.Id},
		{Dice: &Dice{Name: FILLED}, PlayerId: p1.Id}}
	handleQuestionmark(&dr, 0, p1)

	if (*p1.BonusCard)[DOT].CurrentChecks != 0 {
		t.Errorf("Expected bonus to be %d, got %d instead.", 0, (*p1.BonusCard)[DOT].CurrentChecks)
	}
	if dr[0].PlayerId != p1.Id && !dr[0].IsSelected {
		t.Errorf("Expected pID %d, got %d. Expected is selected %t, got %t", p1.Id, dr[0].PlayerId, true, dr[0].IsSelected)
	}
}
