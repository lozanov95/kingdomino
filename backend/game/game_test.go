package game

import (
	"testing"
)

func TestGetDieAllSides(t *testing.T) {
	p1 := NewMockPlayer([]ClientPayload{})
	p2 := NewMockPlayer([]ClientPayload{})
	gr := NewGame(p1, p2)
	dr := gr.GetDieAllSides(0)

	if dr[0].Dice.Name != QUESTIONMARK {
		t.Errorf("Expected %s, got %s", QUESTIONMARK.String(), dr[0].Dice.Name)
	}
	if dr[1].Dice.Name != DOT {
		t.Errorf("Expected %s, got %s", DOT.String(), dr[1].Dice.Name)
	}
	if dr[2].Dice.Name != DOUBLELINE {
		t.Errorf("Expected %s, got %s", DOUBLELINE.String(), dr[2].Dice.Name)
	}
	if dr[3].Dice.Name != LINE {
		t.Errorf("Expected %s, got %s", QUESTIONMARK.String(), dr[3].Dice.Name)
	}
	if dr[4].Dice.Name != DOUBLEDOT {
		t.Errorf("Expected %s, got %s", QUESTIONMARK.String(), dr[4].Dice.Name)
	}
	if dr[5].Dice.Name != DOT {
		t.Errorf("Expected %s, got %s", DOT.String(), dr[5].Dice.Name)
	}
}
