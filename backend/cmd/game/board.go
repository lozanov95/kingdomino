package game

import "fmt"

type Board [5][7]Badge

func NewBoard() *Board {
	var newB Board
	newB[2][3] = Badge{CASTLE, 0}
	return &newB
}

func (b *Board) Print() {
	fmt.Println("BOARD")
	lb := "----------------------------------------------------------"
	fmt.Printf("%s\n", lb)
	var cell Badge
	for i := 0; i < len(b); i++ {
		fmt.Printf("%d| ", i)
		for j := 0; j < len(b[i]); j++ {
			cell = b[i][j]
			fmt.Print(cell.name.String(), cell.nobles, " | ")
		}
		fmt.Printf("\n%s\n", lb)
	}
}
