package game

import "log"

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
	case EMPTY:
		return "EMPTY"
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
		log.Println("Invalid choice")
		return "INVALID"
	}
}
