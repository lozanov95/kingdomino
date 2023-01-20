package server_test

import (
	"testing"

	"github.com/lozanov95/kingdomino/backend/cmd/game"
	"github.com/lozanov95/kingdomino/backend/cmd/server"
	"golang.org/x/net/websocket"
)

func TestJoinRoom(t *testing.T) {
	c1 := &websocket.Conn{}
	c2 := &websocket.Conn{}
	c3 := &websocket.Conn{}
	gr := server.NewGameRoom(make(chan string))
	p1 := game.NewPlayer(c1)
	p2 := game.NewPlayer(c2)
	p3 := game.NewPlayer(c3)

	if err := gr.Join(p1); err != nil {
		t.Error(err)
	}
	if err := gr.Join(p2); err != nil {
		t.Error(err)
	}

	if err := gr.Join(p3); err != server.ErrGameRoomFull {
		t.Error("incorrect error on full room")
	}
	if len(gr.Players) != gr.PlayerLimit {
		t.Errorf("expected %d players, got %d", gr.PlayerLimit, len(gr.Players))
	}
}
