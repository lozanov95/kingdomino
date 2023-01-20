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

func (b *Board) IsThereOccupiedNeighbourCell(row, cell int) bool {
	if b.isCellOccupied(row-1, cell) ||
		b.isCellOccupied(row+1, cell) ||
		b.isCellOccupied(row, cell-1) ||
		b.isCellOccupied(row, cell+1) {
		return true
	}

	return false
}

func (b *Board) IsThereFreeNeighbourCell(row, cell int) bool {
	if b.isCellEmpty(row-1, cell) ||
		b.isCellEmpty(row+1, cell) ||
		b.isCellEmpty(row, cell-1) ||
		b.isCellEmpty(row, cell+1) {
		return true
	}

	return false
}

func (b *Board) isCellEmpty(row, cell int) bool {
	if (row >= len(b) || row < 0) ||
		(cell >= len(b[row]) || cell < 0) {
		return false
	}

	return b[row][cell].Name == EMPTY
}

func (b *Board) isCellOccupied(row, cell int) bool {
	if (row >= len(b) || row < 0) ||
		(cell >= len(b[row]) || cell < 0) {
		return false
	}

	return b[row][cell].Name != EMPTY
}
