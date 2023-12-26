package color

import (
	"fmt"

	"github.com/go-hl/color/background"
	"github.com/go-hl/color/effect"
	"github.com/go-hl/color/foreground"
)

// Print print with fmt.Print() colored.
func Print(
	foreground *foreground.Foreground,
	background *background.Background,
	effects []*effect.Effect,
	args ...any,
) {
	fmt.Print(make(foreground, background, effects, args...))
}

// Println print with fmt.Println() colored.
func Println(
	foreground *foreground.Foreground,
	background *background.Background,
	effects []*effect.Effect,
	args ...any,
) {
	fmt.Println(make(foreground, background, effects, args...))
}

// Printf print with fmt.Printf() colored.
func Printf(
	foreground *foreground.Foreground,
	background *background.Background,
	effects []*effect.Effect,
	format string,
	args ...any,
) {
	fmt.Printf(make(foreground, background, effects, format), args...)
}
