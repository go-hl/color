// Package color is the package for colored terminal output.
package color

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-hl/color/background"
	"github.com/go-hl/color/effect"
	"github.com/go-hl/color/foreground"
	"github.com/go-hl/color/format"
)

func make(
	foreground *foreground.Foreground,
	background *background.Background,
	effects []*effect.Effect,
	args ...any,
) string {
	var converted []string
	for _, arg := range args {
		converted = append(converted, fmt.Sprintf("%v", arg))
	}
	formatter := "\033["
	if foreground != nil {
		formatter += strconv.Itoa(int(*foreground))
	}
	if background != nil {
		formatter += ";" + strconv.Itoa(int(*background))
	}
	for _, effect := range effects {
		formatter += ";" + strconv.Itoa(int(*effect))
	}
	return formatter + "m" + strings.Join(converted, " ") + "\033[m"
}

// NewFormat returns a new pointer of Format struct with your values alredy setupeds.
func NewFormat(
	foreground foreground.Foreground,
	background background.Background,
	effects []effect.Effect,
) format.Format {
	format := format.Format{
		Foreground: foreground,
		Background: background,
		Effects:    effects,
	}
	format.Prepare()
	return format
}
