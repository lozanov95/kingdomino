package game

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Bonus struct {
	RequiredChecks int
	CurrentChecks  int
	Eligible       bool
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

// func (b *Bonus) MarshalJSON() ([]byte, error) {
// 	return []byte(fmt.Sprintf("{\"requiredChecks\":%d,\"currentChecks\":%d,\"eligible\":%t}", b.RequiredChecks, b.CurrentChecks, b.Eligible)), nil
// }

func (bm *BonusMap) MarshalJSON() ([]byte, error) {
	var sb strings.Builder

	sb.WriteString("{\"bonusCard\":[")
	var idx int
	for badge, bonus := range *bm {
		sb.WriteString(fmt.Sprintf("{\"name\":%d,", badge))
		jsonBonus, err := json.Marshal(bonus)
		fmt.Println("bonus:", string(jsonBonus))
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

	sb.WriteString("]}")

	return []byte(sb.String()), nil
}
