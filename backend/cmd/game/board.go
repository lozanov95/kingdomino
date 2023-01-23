package game

import (
	"encoding/json"
	"log"
	"strings"
)

type Board [5][7]Badge

type BoardResult struct {
	Count  int
	Nobles int
}

func (br *BoardResult) Add(newBR BoardResult) {
	br.Count += newBR.Count
	br.Nobles += newBR.Nobles
}

func (br *BoardResult) CalculatePoints() int {
	return br.Count * br.Nobles
}

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

func (b *Board) CalculateBadgePoints(bt BadgeName) int {
	points := 0
	visited := make(map[DiePos]bool)
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			dp := DiePos{Row: i, Cell: j}
			if !visited[dp] {
				visited[dp] = true
				br := BoardResult{}
				if b.isCellMatching(dp, bt) {
					br.Count++
					br.Nobles += b[i][j].Nobles
				}
				br.Add(b.findNeighbour(dp, bt, visited))
				points += br.CalculatePoints()
			}
		}
	}

	return points
}

func (b Board) findNeighbour(dp DiePos, bt BadgeName, visited map[DiePos]bool) BoardResult {
	br := BoardResult{}
	pos := DiePos{Row: dp.Row + 1, Cell: dp.Cell}
	if !visited[pos] {
		visited[pos] = true
		if b.isCellMatching(pos, bt) {
			br.Count++
			br.Nobles += b[pos.Row][pos.Cell].Nobles
			br.Add(b.findNeighbour(pos, bt, visited))
		}
	}
	pos = DiePos{Row: dp.Row - 1, Cell: dp.Cell}
	if !visited[pos] {
		visited[pos] = true
		if b.isCellMatching(pos, bt) {
			br.Count++
			br.Nobles += b[pos.Row][pos.Cell].Nobles
			br.Add(b.findNeighbour(pos, bt, visited))
		}
	}
	pos = DiePos{Row: dp.Row, Cell: dp.Cell + 1}
	if !visited[pos] {
		visited[pos] = true
		if b.isCellMatching(pos, bt) {
			br.Count++
			br.Nobles += b[pos.Row][pos.Cell].Nobles
			br.Add(b.findNeighbour(pos, bt, visited))
		}
	}
	pos = DiePos{Row: dp.Row, Cell: dp.Cell - 1}
	if !visited[pos] {
		visited[pos] = true
		if b.isCellMatching(pos, bt) {
			br.Count++
			br.Nobles += b[pos.Row][pos.Cell].Nobles
			br.Add(b.findNeighbour(pos, bt, visited))
		}
	}

	return br
}

func (b Board) isCellMatching(dp DiePos, bn BadgeName) bool {
	if dp.Row < 0 || dp.Row >= len(b) || dp.Cell < 0 || dp.Cell >= len(b[0]) {
		return false
	}

	return bn == b[dp.Row][dp.Cell].Name
}
