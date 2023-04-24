package game

type BadgeName uint8

type Dice struct {
	Name   BadgeName `json:"name"`
	Nobles int       `json:"nobles"`
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
