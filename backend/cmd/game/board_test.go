package game

import "testing"

func TestIsCellMatching(t *testing.T) {
	board := Board{
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: LINE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: CASTLE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
	}

	if !board.isCellMatching(DiePos{Row: 0, Cell: 0}, DOT) {
		t.Error("expected the cell to be matching")
	}
	if board.isCellMatching(DiePos{Row: 1, Cell: 0}, DOT) {
		t.Error("expected the cell to not be matching")
	}
}

func TestCalculateBadgePoints(t *testing.T) {
	board := Board{
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: CHECKED, Nobles: 0}, {Name: DOUBLEDOT, Nobles: 1}, {Name: DOUBLEDOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: CHECKED, Nobles: 0}, {Name: CHECKED, Nobles: 1}, {Name: CHECKED, Nobles: 0}, {Name: DOUBLEDOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: DOT, Nobles: 0}, {Name: FILLED, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: CASTLE, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: FILLED, Nobles: 2}, {Name: FILLED, Nobles: 0}, {Name: FILLED, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOUBLEDOT, Nobles: 1}, {Name: DOUBLEDOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
		[7]Badge{{Name: FILLED, Nobles: 1}, {Name: FILLED, Nobles: 1}, {Name: FILLED, Nobles: 0}, {Name: DOT, Nobles: 0}, {Name: DOUBLEDOT, Nobles: 1}, {Name: DOT, Nobles: 0}, {Name: DOT, Nobles: 0}},
	}

	pts, dms := board.CalculateBadgePoints(CHECKED)
	if pts != 4 {
		t.Errorf("expected %d points, got %d", 4, pts)
	}
	if dms != 1 {
		t.Errorf("expected %d domains, got %d", 1, dms)
	}
	pts, dms = board.CalculateBadgePoints(FILLED)
	if pts != 28 {
		t.Errorf("expected %d points, got %d", 28, pts)
	}
	if dms != 1 {
		t.Errorf("expected %d domains, got %d", 1, dms)
	}
	pts, dms = board.CalculateBadgePoints(LINE)
	if pts != 0 {
		t.Errorf("expected %d points, got %d", 0, pts)
	}
	if dms != 0 {
		t.Errorf("expected %d domains, got %d", 0, dms)
	}
	pts, dms = board.CalculateBadgePoints(DOUBLEDOT)
	if pts != 9 {
		t.Errorf("expected %d points, got %d", 9, pts)
	}
	if dms != 2 {
		t.Errorf("expected %d domains, got %d", 2, dms)
	}
}
