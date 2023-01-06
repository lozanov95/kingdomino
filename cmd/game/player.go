package game

type Player struct {
	name      string
	board     *Board
	bonuscard *map[BadgeName]Bonus
}

// Creates a new player instance and returns a pointer to it.
func NewPlayer(name string) *Player {
	return &Player{
		name:      name,
		board:     NewBoard(),
		bonuscard: NewBonusMap(),
	}
}

// Increases the bonus of a specific card
func (p *Player) IncreaseBonus(b Badge) {
	if b.nobles != 0 {
		return
	}

	tmp := (*p.bonuscard)[b.name]
	tmp.Increment()
	(*p.bonuscard)[b.name] = tmp
}
