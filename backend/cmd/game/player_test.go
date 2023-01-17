package game

import (
	"testing"

	"golang.org/x/net/websocket"
)

func TestIncrementWithoutNobles(t *testing.T) {
	b := Badge{name: LINE, nobles: 0}
	p := NewPlayer([]byte("test"), &websocket.Conn{})

	p.IncreaseBonus(b)

	if (*p.bonuscard)[LINE].CurrentChecks != 1 {
		t.Errorf("Expected bonus of %s to be 1, but it is %d", LINE.String(), (*p.bonuscard)[LINE].CurrentChecks)
	}
}

func TestIncrementWithNobles(t *testing.T) {
	b := Badge{name: LINE, nobles: 1}
	p := NewPlayer([]byte("test"), &websocket.Conn{})

	p.IncreaseBonus(b)

	if (*p.bonuscard)[LINE].CurrentChecks != 0 {
		t.Errorf("Expected bonus of %s to be 0, but it is %d", LINE.String(), (*p.bonuscard)[LINE].CurrentChecks)
	}
}
