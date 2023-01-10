package game

type Bonus struct {
	RequiredChecks int
	CurrentChecks  int
}

func NewBonusMap() *map[BadgeName]Bonus {
	return &map[BadgeName]Bonus{
		DOT:        {RequiredChecks: 5},
		LINE:       {RequiredChecks: 5},
		DOUBLEDOT:  {RequiredChecks: 4},
		DOUBLELINE: {RequiredChecks: 4},
		FILLED:     {RequiredChecks: 3},
		CHECKED:    {RequiredChecks: 3},
	}
}

func (b *Bonus) IsCompleted() bool {
	return b.CurrentChecks == b.RequiredChecks
}

func (b *Bonus) Increment() {
	if !b.IsCompleted() {
		b.CurrentChecks++
	}
}
