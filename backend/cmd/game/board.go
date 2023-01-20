package game

import (
	"encoding/json"
	"log"
	"strings"
)

type Board [5][7]Badge

func NewBoard() *Board {
	var newB Board
	newB[2][3] = Badge{CASTLE, 0}
	return &newB
}

func (b *Board) Json() ([]byte, error) {
	res, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

func (b *Board) MarshalJSON() ([]byte, error) {
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < len(b); i++ {
		sb.WriteString("[")
		for j := 0; j < len(b[i]); j++ {
			badge, err := json.Marshal(b[i][j])
			if err != nil {
				log.Println(err)
				return nil, err
			}
			sb.Write(badge)
			if j < len(b[i])-1 {
				sb.WriteString(",")
			}
		}
		sb.WriteString("]")
		if i < len(b)-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString("]")
	return []byte(sb.String()), nil
}

func (b *Board) IsValidPlacementPos(row, cell int) bool {
	if row >= len(b) || row < 0 ||
		cell >= len(b[row]) || cell < 0 ||
		b[row][cell].Name != EMPTY {
		return false
	}
	startI := row - 1
	startJ := cell - 1
	if startI < 0 {
		startI = 0
	}
	if startJ < 0 {
		startJ = 0
	}
	maxI := row + 1
	maxJ := cell + 1
	if maxI > 4 {
		maxI = 4
	}
	if maxJ > 6 {
		maxJ = 6
	}

	for i := startI; i <= maxI; i++ {
		for j := startJ; j <= maxJ; j++ {
			if b[i][j].Name != EMPTY && (i != row || j != cell) {
				return true
			}
		}
	}

	return false
}

func (b *Board) IsThereFreeNeighbourCell(row, cell int) bool {
	if b.doesCellMatchBadge(EMPTY, row-1, cell) ||
		b.doesCellMatchBadge(EMPTY, row+1, cell) ||
		b.doesCellMatchBadge(EMPTY, row, cell-1) ||
		b.doesCellMatchBadge(EMPTY, row, cell+1) {
		log.Printf("Cell [%d][%d] has a valid neighbour", row, cell)
		return true
	}

	return false
}

func (b *Board) doesCellMatchBadge(bn BadgeName, row, cell int) bool {
	if (row >= len(b) || row < 0) ||
		(cell >= len(b[row]) || cell < 0) {
		return false
	}

	return b[row][cell].Name == bn
}
