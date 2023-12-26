package effect

// Effect is the Effect type.
type Effect int

// Effect consts.
const (
	None       Effect = -1
	Bold       Effect = 1
	Underlined Effect = 4
	Flashing   Effect = 5
	Reverse    Effect = 7
)

// Effects is a array with all effect types.
var Effects = [...]Effect{
	None,
	Bold,
	Underlined,
	Flashing,
	Reverse,
}

// Index conts.
const (
	NoneIndex = iota
	BoldIndex
	UnderlinedIndex
	FlashingIndex
	ReverseIndex
)
