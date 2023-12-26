package format

import (
	"reflect"
	"strconv"

	"github.com/go-hl/color/background"
	"github.com/go-hl/color/effect"
	"github.com/go-hl/color/foreground"
)

// Colors constants.
const (
	Break    = `%%<`
	Continue = `%%>`
	reset    = "\033[m"
)

func validate[Type foreground.Foreground | background.Background | effect.Effect](format any, types []Type) (found bool) {
	for _, Type := range types {
		if format != Type {
			continue
		} else {
			found = true
		}
	}
	return
}

func convert[Type foreground.Foreground | background.Background | effect.Effect](format Type) (string, bool) {
	switch reflect.TypeOf(format).Name() {
	case "Foreground":
		if !validate(format, foreground.Foregrounds[:]) {
			panic("wrong or missing foreground type")
		}
		if format == Type(foreground.None) {
			return "", true
		}
	case "Background":
		if !validate(format, background.Backgrounds[:]) {
			panic("wrong or missing background type")
		}
		if format == Type(background.None) {
			return "", true
		}
	case "Effect":
		if !validate(format, effect.Effects[:]) {
			panic("wrong or missing effect type")
		}
		if format == Type(effect.None) {
			return "", true
		}
	}
	return strconv.Itoa(int(format)), false
}
