package game

type Game struct {
	dices [4]Dice
	p1    *Player
	p2    *Player
}

// Starts a new game and returns it as an instance
func NewGame(p1Name, p2Name string) *Game {
	g := Game{
		p1: NewPlayer(p1Name),
		p2: NewPlayer(p2Name),
	}
	g.setupDice()
	return &g
}

// Creates and setups the correct Dice sides
func (g *Game) setupDice() {
	g.dices[0] = Dice{
		[6]Badge{
			{name: QUESTIONMARK, nobles: 0},
			{name: DOT, nobles: 0},
			{name: DOUBLELINE, nobles: 1},
			{name: LINE, nobles: 0},
			{name: DOUBLEDOT, nobles: 1},
			{name: DOT, nobles: 0},
		},
	}
	g.dices[1] = Dice{
		[6]Badge{
			{name: CHECKED, nobles: 0},
			{name: LINE, nobles: 0},
			{name: DOUBLEDOT, nobles: 1},
			{name: DOUBLELINE, nobles: 1},
			{name: DOT, nobles: 0},
			{name: FILLED, nobles: 0},
		},
	}
	g.dices[2] = Dice{
		[6]Badge{
			{name: DOUBLEDOT, nobles: 0},
			{name: LINE, nobles: 0},
			{name: FILLED, nobles: 2},
			{name: CHECKED, nobles: 2},
			{name: DOT, nobles: 0},
			{name: DOUBLELINE, nobles: 0},
		},
	}
	g.dices[3] = Dice{
		[6]Badge{
			{name: DOUBLELINE, nobles: 0},
			{name: LINE, nobles: 1},
			{name: QUESTIONMARK, nobles: 0},
			{name: LINE, nobles: 0},
			{name: DOT, nobles: 1},
			{name: DOUBLEDOT, nobles: 0},
		},
	}
}

func (g *Game) RollDice() [4]Badge {
	d := [4]Badge{}
	for i, r := range g.dices {
		d[i] = r.Roll()
	}

	return d
}
