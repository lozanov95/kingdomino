package game

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Bonus struct {
	RequiredChecks int  `json:"requiredChecks"`
	CurrentChecks  int  `json:"currentChecks"`
	Eligible       bool `json:"eligible"`
}

type BonusMap map[BadgeName]Bonus
type PowerType int64

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

func (b *Bonus) IsCompleted() bool {
	return b.CurrentChecks == b.RequiredChecks
}

func (b *Bonus) Increment() {
	if !b.IsCompleted() && b.Eligible {
		b.CurrentChecks++
	}

	if b.IsCompleted() {
		b.Eligible = false
	}
}

func (bm *BonusMap) MarshalJSON() ([]byte, error) {
	var sb strings.Builder
	sb.WriteString("[")
	var idx int
	for badge, bonus := range *bm {
		sb.WriteString(fmt.Sprintf("{\"name\":%d,", badge))
		jsonBonus, err := json.Marshal(bonus)

		if err != nil {
			return nil, err
		}
		jb := strings.Replace(string(jsonBonus), "{", "", 1)
		sb.WriteString(jb)
		if idx < len(*bm)-1 {
			sb.WriteString(",")
		}
		idx++
	}

	sb.WriteString("]")

	return []byte(sb.String()), nil
}

func (bm *BonusMap) AddBonus(d Badge) {
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
		return DOUBLEDOT
	case PWRDomainPoints:
		return FILLED
	case PWRAddNoble:
		return CHECKED
	default:
		return EMPTY
	}
}

// Returns if the bonus for a specific badge is completed
func (bm *BonusMap) IsBonusCompleted(pt PowerType) bool {
	bonus := (*bm)[getBonusBadge(pt)]
	return bonus.IsCompleted()
}
