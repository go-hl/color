package background

// Background is the type fo background.
type Background int

// Background consts.
const (
	None Background = iota + 39
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Backgrounds is a array with all background types.
var Backgrounds = [...]Background{
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
