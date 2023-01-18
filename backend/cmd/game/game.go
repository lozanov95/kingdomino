package game

type Game struct {
	dices [4]Dice
	p1    *Player
	p2    *Player
}

// Starts a new game and returns it as an instance
func NewGame(p1, p2 *Player) *Game {
	g := Game{
		p1: p1,
		p2: p2,
	}
	g.setupDice()
	return &g
}

// Creates and setups the correct Dice sides
func (g *Game) setupDice() {
	g.dices[0] = Dice{
		[6]Badge{
			{Name: QUESTIONMARK, Nobles: 0},
			{Name: DOT, Nobles: 0},
			{Name: DOUBLELINE, Nobles: 1},
			{Name: LINE, Nobles: 0},
			{Name: DOUBLEDOT, Nobles: 1},
			{Name: DOT, Nobles: 0},
		},
	}
	g.dices[1] = Dice{
		[6]Badge{
			{Name: CHECKED, Nobles: 0},
			{Name: LINE, Nobles: 0},
			{Name: DOUBLEDOT, Nobles: 1},
			{Name: DOUBLELINE, Nobles: 1},
			{Name: DOT, Nobles: 0},
			{Name: FILLED, Nobles: 0},
		},
	}
	g.dices[2] = Dice{
		[6]Badge{
			{Name: DOUBLEDOT, Nobles: 0},
			{Name: LINE, Nobles: 0},
			{Name: FILLED, Nobles: 2},
			{Name: CHECKED, Nobles: 2},
			{Name: DOT, Nobles: 0},
			{Name: DOUBLELINE, Nobles: 0},
		},
	}
	g.dices[3] = Dice{
		[6]Badge{
			{Name: DOUBLELINE, Nobles: 0},
			{Name: LINE, Nobles: 1},
			{Name: QUESTIONMARK, Nobles: 0},
			{Name: LINE, Nobles: 0},
			{Name: DOT, Nobles: 1},
			{Name: DOUBLEDOT, Nobles: 0},
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
