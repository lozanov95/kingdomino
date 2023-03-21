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

type BoardPlacementInput struct {
	// Provide the previous placement position if this will be the 2nd die placement
	PrevPosition DiePos

	// Pointer to the user's board
	Board *Board

	// Allows the user to place dice separately
	SeparateDice bool

	// Follow Connection Rules
	IgnoreConnectionRules bool
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

// Calculate the points for a given badge
func (b *Board) CalculateBadgePoints(bt BadgeName) (points, domains int) {
	points = 0
	domains = 0
	visited := make(map[DiePos]bool)
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			dp := DiePos{Row: i, Cell: j}
			if visited[dp] {
				continue
			}

			visited[dp] = true
			if !b.isCellMatching(dp, bt) {
				continue
			}

			br := BoardResult{}
			br.Count++
			br.Nobles += b[i][j].Nobles
			domains++
			br.Add(b.findNeighbour(dp, bt, visited))
			points += br.CalculatePoints()
		}
	}

	return points, domains
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

// Checks if the cell has a specific badge
func (b Board) isCellMatching(dp DiePos, bn BadgeName) bool {
	if dp.Row < 0 || dp.Row >= len(b) || dp.Cell < 0 || dp.Cell >= len(b[0]) {
		return false
	}

	return bn == b[dp.Row][dp.Cell].Name
}

// Checks if you can place a domino.
// The selected cell must be free. If this is the 1st badge of the domino, there must be a free neighbouring cell
// If this is the 2nd badge, the position must be free and it should be a neighbour of the first domino
func (bpi *BoardPlacementInput) IsValid(newPos *DiePos) bool {
	if bpi.SeparateDice {
		return bpi.Board[newPos.Row][newPos.Cell].Name == EMPTY
	}

	emptyPos := DiePos{}
	if bpi.PrevPosition == emptyPos {
		return bpi.Board[newPos.Row][newPos.Cell].Name == EMPTY && bpi.Board.IsThereFreeNeighbourCell(newPos.Row, newPos.Cell)
	}

	if (newPos.Row-1 == bpi.PrevPosition.Row && newPos.Cell == bpi.PrevPosition.Cell) ||
		(newPos.Row+1 == bpi.PrevPosition.Row && newPos.Cell == bpi.PrevPosition.Cell) ||
		(newPos.Row == bpi.PrevPosition.Row && newPos.Cell-1 == bpi.PrevPosition.Cell) ||
		(newPos.Row == bpi.PrevPosition.Row && newPos.Cell+1 == bpi.PrevPosition.Cell) {
		return true
	}

	return false
}
