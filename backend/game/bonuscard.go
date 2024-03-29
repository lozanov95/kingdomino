package game

import "encoding/json"

type Bonus struct {
	RequiredChecks int  `json:"requiredChecks"`
	CurrentChecks  int  `json:"currentChecks"`
	Eligible       bool `json:"eligible"`
	Used           bool `json:"used"`
}

type PlayerPower struct {
	// The power type
	Type PowerType `json:"type"`

	// User friendly description of the power
	Description string `json:"description"`

	// Does the player wants to use the power
	Use bool `json:"use"`

	// Does the player answered the use power prompt
	Confirmed bool `json:"confirmed"`
}

type BonusMap map[BadgeName]Bonus
type PowerType uint8

type BonusResult struct {
	BadgeName `json:"name"`
	Bonus
}

const (
	// Grants no wizard power
	PWRNoPower PowerType = iota

	//You can play your domino without following
	// the Connection Rules.
	PWRNoConnectionRules

	// You can separate your dice to fill in your map.
	// Each die must respect the Connection Rules.
	PWRSeparateDominos

	// During a turn where you are the player A,
	// you can immediately pick your 2 dice.
	PWRPickTwoDice

	// You can turn one of the dice in your
	// domino around so that it shows any face
	// of your choice.
	PWRSelectDieSideOfChoice

	// Choose a coat of arms. Each different
	// DOMAIN with this coat of arms will earn you
	// 3 prestige points at the end of the game.
	PWRDomainPoints

	// Add one cross to the coat of arms of your
	// choosing.
	PWRAddNoble
)

func NewBonusMap() *BonusMap {
	return &BonusMap{
		DOT:        {RequiredChecks: 5, Eligible: true},
		LINE:       {RequiredChecks: 5, Eligible: true},
		DOUBLEDOT:  {RequiredChecks: 4, Eligible: true},
		DOUBLELINE: {RequiredChecks: 4, Eligible: true},
		FILLED:     {RequiredChecks: 3, Eligible: true},
		CHECKED:    {RequiredChecks: 3, Eligible: true},
	}
}

// Returns true if you have collected enough badges
func (b *Bonus) IsCompleted() bool {
	return b.CurrentChecks == b.RequiredChecks
}

func (b *Bonus) IsUsable() bool {
	return b.IsCompleted() && !b.Used
}

func (b *Bonus) Increment() {
	if !b.IsCompleted() && b.Eligible {
		b.CurrentChecks++
	}

	if b.IsCompleted() {
		b.Eligible = false
	}
}

func (bm *BonusMap) MarkUsed(pt PowerType) {
	bonus := (*bm)[getBonusBadge(pt)]
	bonus.Used = true
	(*bm)[getBonusBadge(pt)] = bonus
}

func (bm *BonusMap) MarshalJSON() ([]byte, error) {
	var bonuses []BonusResult

	for bn, b := range *bm {
		bonuses = append(bonuses, BonusResult{BadgeName: bn, Bonus: b})
	}

	return json.Marshal(bonuses)
}

func (bm *BonusMap) AddBonus(d Dice) {
	if d.Nobles > 0 {
		return
	}

	b := (*bm)[d.Name]
	b.Increment()
	(*bm)[d.Name] = b
}

func (bm *BonusMap) IsThereACompletedBonus() bool {
	for _, bonus := range *bm {
		if bonus.IsCompleted() {
			return true
		}
	}

	return false
}

// Returns the badge name that is required for the given power
func getBonusBadge(pt PowerType) BadgeName {
	switch pt {
	case PWRNoConnectionRules:
		return DOT
	case PWRSeparateDominos:
		return LINE
	case PWRPickTwoDice:
		return DOUBLEDOT
	case PWRSelectDieSideOfChoice:
		return DOUBLELINE
	case PWRDomainPoints:
		return FILLED
	case PWRAddNoble:
		return CHECKED
	default:
		return EMPTY
	}
}

// Returns the PowerType for the given badge
func getBonusType(b BadgeName) PowerType {
	switch b {
	case DOT:
		return PWRNoConnectionRules
	case LINE:
		return PWRSeparateDominos
	case DOUBLEDOT:
		return PWRPickTwoDice
	case DOUBLELINE:
		return PWRSelectDieSideOfChoice
	case FILLED:
		return PWRDomainPoints
	case CHECKED:
		return PWRAddNoble
	default:
		return PWRNoPower
	}
}

// Returns if the bonus for a specific badge is completed
func (bm *BonusMap) IsBonusCompleted(pt PowerType) bool {
	bonus := (*bm)[getBonusBadge(pt)]
	return bonus.IsCompleted()
}

// Returns if the bonus is completed and haven't been used yet
func (bm *BonusMap) IsBonusUsable(pt PowerType) bool {
	bonus := (*bm)[getBonusBadge(pt)]
	return bonus.IsUsable()
}
