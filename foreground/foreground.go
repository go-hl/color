package foreground

// Foreground is the foreground type.
type Foreground int

// Foreground consts.
const (
	None Foreground = iota + 29
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Foregrounds is a array with all foreground types.
var Foregrounds = [...]Foreground{
	None,
	Black,
	Red,
	Green,
	Yellow,
	Blue,
	Magenta,
	Cyan,
	White,
}

// Index consts.
const (
	NoneIndex = iota
	BlackIndex
	RedIndex
	GreenIndex
	YellowIndex
	BlueIndex
	MagentaIndex
	CyanIndex
	WhiteIndex
)
