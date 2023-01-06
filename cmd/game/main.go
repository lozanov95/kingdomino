package game

import "fmt"

func main() {
	g := NewGame("Vasko", "Didi")
	g.p1.board.Print()
	d := g.RollDice()
	for _, die := range d {
		fmt.Printf("%s %d\n", die.name.String(), die.nobles)
	}
	fmt.Println(g.p1.name, g.p2.name)
}
