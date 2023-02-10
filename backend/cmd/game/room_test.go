package game

import (
	"testing"

	"golang.org/x/net/websocket"
)

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
