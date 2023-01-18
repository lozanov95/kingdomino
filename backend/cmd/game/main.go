package game

import (
	"fmt"
)

func TestGame() {
	g := NewGame(&Player{}, &Player{})
	g.p1.Board.Print()
	d := g.RollDice()
	for _, die := range d {
		fmt.Printf("%s %d\n", die.name.String(), die.nobles)
	}
}
