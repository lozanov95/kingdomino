package game

import (
	"fmt"
)

type BadgeName int

type Badge struct {
	name   BadgeName
	nobles int
}

const (
	EMPTY BadgeName = iota
	CASTLE
	DOT
	LINE
	DOUBLEDOT
	DOUBLELINE
	FILLED
	CHECKED
	QUESTIONMARK
)

// Returns the string representation of the BADGE
func (b BadgeName) String() string {
	switch b {
	case CASTLE:
		return "CASTLE"
	case DOT:
		return "DOT"
	case LINE:
		return "LINE"
	case DOUBLEDOT:
		return "DOUBLEDOT"
	case DOUBLELINE:
		return "DOUBLELINE"
	case FILLED:
		return "FILLED"
	case CHECKED:
		return "CHECKED"
	case QUESTIONMARK:
		return "QUESTIONMARK"
	default:
		return "EMPTY"
	}
}

func (b Badge) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"name\":%d,\"nobles\":%d}", b.name, b.nobles)), nil
}
