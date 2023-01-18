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
	if !b.IsCompleted() {
		b.CurrentChecks++
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
