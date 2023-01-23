package game

import (
	"testing"

	"golang.org/x/net/websocket"
)

func TestIncrementWithoutNobles(t *testing.T) {
	b := Badge{Name: LINE, Nobles: 0}
	p := NewPlayer(&websocket.Conn{})

	p.IncreaseBonus(b)

	if (*p.BonusCard)[LINE].CurrentChecks != 1 {
		t.Errorf("Expected bonus of %s to be 1, but it is %d", LINE.String(), (*p.BonusCard)[LINE].CurrentChecks)
	}
}

func TestIncrementWithNobles(t *testing.T) {
	b := Badge{Name: LINE, Nobles: 1}
	p := NewPlayer(&websocket.Conn{})

	p.IncreaseBonus(b)

	if (*p.BonusCard)[LINE].CurrentChecks != 0 {
		t.Errorf("Expected bonus of %s to be 0, but it is %d", LINE.String(), (*p.BonusCard)[LINE].CurrentChecks)
	}
}

func TestCalculatePoints(t *testing.T) {
	p := NewPlayer(&websocket.Conn{})
	p.Board = &Board{
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: CASTLE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
	}

	score := p.CalculateScore()
	if score != 0 {
		t.Errorf("expected %d, got %d", 0, score)
	}

	p1 := NewPlayer(&websocket.Conn{})
	p1.Board = &Board{
		[7]Badge{{Name: LINE, Nobles: 1}, {Name: DOT, Nobles: 0}, {Name: CHECKED, Nobles: 2}, {Name: LINE, Nobles: 1}, {Name: LINE, Nobles: 0}, {Name: FILLED, Nobles: 1}, {Name: FILLED, Nobles: 2}},
		[7]Badge{{Name: LINE, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: CHECKED, Nobles: 0}, {Name: CHECKED, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: FILLED, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: LINE, Nobles: 1}, {Name: LINE, Nobles: 2}, {Name: CHECKED, Nobles: 1}, {Name: CASTLE, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: LINE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: CHECKED, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
	}

	score = p1.CalculateScore()
	if score != 55 {
		t.Errorf("expected %d, got %d", 55, score)
	}

}
