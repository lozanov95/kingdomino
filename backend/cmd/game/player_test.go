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
	p1.SelectedCoatOfArms = LINE
	p1.Board = &Board{
		[7]Badge{{Name: LINE, Nobles: 1}, {Name: DOT, Nobles: 0}, {Name: CHECKED, Nobles: 2}, {Name: LINE, Nobles: 1}, {Name: LINE, Nobles: 0}, {Name: FILLED, Nobles: 1}, {Name: FILLED, Nobles: 2}},
		[7]Badge{{Name: LINE, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: CHECKED, Nobles: 0}, {Name: CHECKED, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: FILLED, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: LINE, Nobles: 1}, {Name: LINE, Nobles: 2}, {Name: CHECKED, Nobles: 1}, {Name: CASTLE, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: LINE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: CHECKED, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: LINE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
	}

	score = p1.CalculateScore()
	if score != 64 {
		t.Errorf("expected %d, got %d", 64, score)
	}
}

func TestPlaceDomino(t *testing.T) {
	player := NewMockPlayer([]ClientPayload{
		{SelectedDie: 0}, {DiePos: DiePos{Cell: 3, Row: 3}},
		{SelectedDie: 0}, {DiePos: DiePos{Cell: 4, Row: 3}},
	})
	dice := []Badge{{Name: DOT}, {Name: FILLED}}
	player.Dices = dice
	player.PlaceDomino(&dice)

	row, cell := 3, 3
	if player.Board[row][cell].Name != DOT {
		t.Errorf("Expected board[%d][%d] to be %s, got %s", row, cell, DOT.String(), player.Board[row][cell].Name.String())
	}
	row, cell = 3, 4
	if player.Board[row][cell].Name != FILLED {
		t.Errorf("Expected board[%d][%d] to be %s, got %s", row, cell, FILLED.String(), player.Board[row][cell].Name.String())
	}
}

func TestPlaceDominoInvalidInput(t *testing.T) {
	player := NewMockPlayer([]ClientPayload{
		{SelectedDie: 0}, {DiePos: DiePos{Cell: 0, Row: 0}},
		{SelectedDie: 0}, {DiePos: DiePos{Cell: 3, Row: 3}},
		{SelectedDie: 0}, {DiePos: DiePos{Cell: 4, Row: 3}},
	})
	dice := []Badge{{Name: DOT}, {Name: FILLED}}
	player.Dices = dice
	player.PlaceDomino(&dice)

	row, cell := 3, 3
	if player.Board[row][cell].Name != DOT {
		t.Errorf("Expected board[%d][%d] to be %s, got %s", row, cell, DOT.String(), player.Board[row][cell].Name.String())
	}
	row, cell = 3, 4
	if player.Board[row][cell].Name != FILLED {
		t.Errorf("Expected board[%d][%d] to be %s, got %s", row, cell, FILLED.String(), player.Board[row][cell].Name.String())
	}
}

func TestSeparatedDomino(t *testing.T) {
	player := NewMockPlayer([]ClientPayload{
		{SelectedDie: 0}, {DiePos: DiePos{Row: 0, Cell: 0}},
		{SelectedDie: 0}, {DiePos: DiePos{Row: 4, Cell: 6}},
	})
	player.Board = &Board{
		{{Name: EMPTY}, {Name: LINE}, {Name: LINE}, {Name: DOT}, {Name: DOT}, {Name: DOT}, {Name: DOT}},
		{{Name: DOT}, {Name: LINE}, {Name: LINE}, {Name: DOT}, {Name: DOT}, {Name: DOT}, {Name: DOT}},
		{{Name: DOT}, {Name: EMPTY}, {Name: DOT}, {Name: CASTLE}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: DOT}, {Name: LINE}, {Name: LINE}, {Name: DOT}, {Name: DOT}, {Name: DOT}, {Name: DOT}},
		{{Name: DOT}, {Name: LINE}, {Name: LINE}, {Name: DOT}, {Name: DOT}, {Name: DOT}, {Name: EMPTY}},
	}
	dice := []Badge{{Name: DOT}, {Name: FILLED}}
	player.Dices = dice
	player.PlaceSeparatedDomino(&dice)
	row, cell := 0, 0
	if player.Board[row][cell].Name != DOT {
		t.Errorf("Expected board[%d][%d] to be %s, got %s", row, cell, DOT.String(), player.Board[row][cell].Name.String())
	}

	row, cell = 4, 6
	if player.Board[row][cell].Name != FILLED {
		t.Errorf("Expected board[%d][%d] to be %s, got %s", row, cell, FILLED.String(), player.Board[row][cell].Name.String())
	}
}

func TestUseAddNoblePower(t *testing.T) {
	p := NewMockPlayer([]ClientPayload{
		{DiePos: DiePos{Row: 2, Cell: 2}},
	})
	p.AddBonus(Badge{Name: CHECKED})
	p.AddBonus(Badge{Name: CHECKED})
	p.AddBonus(Badge{Name: CHECKED})
	testBoard := Board{
		{{Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: FILLED}, {Name: FILLED}, {Name: FILLED}, {Name: CASTLE}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
	}
	p.Board = &(testBoard)
	p.handleUseAddNoblePower()

	testBadge := Badge{Name: FILLED, Nobles: 1}
	if p.Board[2][2] != testBadge {
		t.Errorf("Expected %+v, got %+v", testBadge, p.Board[2][2])
	}
	if p.IsBonusUsable(PWRAddNoble) {
		t.Errorf("The PWRAddNoble is not marked as used.")
	}
}
func TestUseAddNoblePowerEmptyValidation(t *testing.T) {
	p := NewMockPlayer([]ClientPayload{
		{DiePos: DiePos{Row: 0, Cell: 0}},
		{DiePos: DiePos{Row: 2, Cell: 2}},
	})
	p.AddBonus(Badge{Name: CHECKED})
	p.AddBonus(Badge{Name: CHECKED})
	p.AddBonus(Badge{Name: CHECKED})
	testBoard := Board{
		{{Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: FILLED}, {Name: FILLED}, {Name: FILLED}, {Name: CASTLE}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
	}
	p.Board = &(testBoard)
	p.handleUseAddNoblePower()

	testBadge := Badge{Name: FILLED, Nobles: 1}
	if p.Board[2][2] != testBadge {
		t.Errorf("Expected %+v, got %+v", testBadge, p.Board[2][2])
	}
}
func TestUseAddNoblePowerNotCompletedBonus(t *testing.T) {
	p := NewMockPlayer([]ClientPayload{
		{DiePos: DiePos{Row: 2, Cell: 2}},
	})
	testBoard := Board{
		{{Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: FILLED}, {Name: FILLED}, {Name: FILLED}, {Name: CASTLE}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
		{{Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}, {Name: EMPTY}},
	}
	p.Board = &(testBoard)
	p.handleUseAddNoblePower()

	testBadge := Badge{Name: FILLED, Nobles: 0}
	if p.Board[2][2] != testBadge {
		t.Errorf("Expected %+v, got %+v", testBadge, p.Board[2][2])
	}
}

func TestGetPlayerPowerChoice(t *testing.T) {
	p := NewMockPlayer([]ClientPayload{
		{PlayerPower: PlayerPower{Use: true}},
		{PlayerPower: PlayerPower{Use: false, Confirmed: true}},
	})
	choice := p.GetPlayerPowerChoice()
	if choice == true {
		t.Errorf("Expected choice to be %t, got %t instead", false, choice)
	}
}

func TestHandleIgnoreConnectionRulesPower(t *testing.T) {
	p := NewMockPlayer([]ClientPayload{{PlayerPower: PlayerPower{Confirmed: true, Use: true}}})
	badge := Badge{Name: getBonusBadge(PWRNoConnectionRules)}
	for i := 0; i < 5; i++ {
		p.IncreaseBonus(badge)
	}

	ignoreRules := p.handleIgnoreConnectionRulesPower()

	if !ignoreRules {
		t.Errorf("Expected ignoreRules to be %t, got %t instead", true, ignoreRules)
	}
	if p.IsBonusUsable(PWRNoConnectionRules) {
		t.Errorf("The power is not marked correctly as unusable")
	}
}
