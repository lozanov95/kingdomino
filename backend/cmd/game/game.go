package game

import (
	"math/rand"
	"time"
)

type Game struct {
	dices [4][]Badge
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
	g.dices = [4][]Badge{
		{
			{Name: QUESTIONMARK, Nobles: 0},
			{Name: DOT, Nobles: 0},
			{Name: DOUBLELINE, Nobles: 1},
			{Name: LINE, Nobles: 0},
			{Name: DOUBLEDOT, Nobles: 1},
			{Name: DOT, Nobles: 0},
		}, {
			{Name: CHECKED, Nobles: 0},
			{Name: LINE, Nobles: 0},
			{Name: DOUBLEDOT, Nobles: 1},
			{Name: DOUBLELINE, Nobles: 1},
			{Name: DOT, Nobles: 0},
			{Name: FILLED, Nobles: 0},
		},
		{
			{Name: DOUBLEDOT, Nobles: 0},
			{Name: LINE, Nobles: 0},
			{Name: FILLED, Nobles: 2},
			{Name: CHECKED, Nobles: 2},
			{Name: DOT, Nobles: 0},
			{Name: DOUBLELINE, Nobles: 0},
		},
		{
			{Name: DOUBLELINE, Nobles: 0},
			{Name: LINE, Nobles: 1},
			{Name: QUESTIONMARK, Nobles: 0},
			{Name: LINE, Nobles: 0},
			{Name: DOT, Nobles: 1},
			{Name: DOUBLEDOT, Nobles: 0},
		},
	}
}

func (g *Game) RollDice() *[]Badge {
	d := make([]Badge, 4)
	for i, r := range g.dices {
		d[i] = Roll(&r)
	}

	return &d
}

// Rolls a die and returns the side it landed on
func Roll(d *[]Badge) Badge {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)
	n := r.Intn(6)

	return (*d)[n]
}
