package game

import (
	"math/rand"
	"time"
)

// Represents a die that has 6 sides
type Dice struct {
	sides [6]Badge
}

// Rolls a die and returns the side it landed on
func (d *Dice) Roll() Badge {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)
	n := r.Intn(6)

	return d.sides[n]
}
