package game

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

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
			fmt.Print(cell.Name.String(), cell.Name, " | ")
		}
		fmt.Printf("\n%s\n", lb)
	}
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
