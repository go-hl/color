package format

import (
	"fmt"
	"strings"

	"github.com/go-hl/color/background"
	"github.com/go-hl/color/effect"
	"github.com/go-hl/color/foreground"
)

// Format struct of format.
type Format struct {
	Foreground foreground.Foreground
	Background background.Background
	Effects    []effect.Effect
	formatter  string
}

// Prepare is auto called on NewFormat() function, case the Format struct is create manually needs call this function manually too.
func (f *Format) Prepare() {
	formatter := "\033["
	if value, none := convert(f.Foreground); !none {
		formatter += value
	}
	if value, none := convert(f.Background); !none {
		formatter += ";" + value
	}
	for _, effect := range f.Effects {
		if value, none := convert(effect); !none {
			formatter += ";" + value
		}
	}
	f.formatter = formatter + "m"
}

func (f Format) parse(args ...any) []any {
	for index, arg := range args {
		if index == 0 {
			if !strings.HasPrefix(fmt.Sprint(arg), Break) {
				arg = f.formatter + fmt.Sprint(arg)
				args[index] = arg
			}
		}
		arg = strings.ReplaceAll(fmt.Sprint(arg), Break, reset)
		args[index] = arg
		arg = strings.ReplaceAll(fmt.Sprint(arg), Continue, f.formatter)
		args[index] = arg
		if index == (len(args) - 1) {
			args[index] = fmt.Sprint(arg) + reset
		}
	}
	return args
}

func (f Format) veirfy() {
	if f.formatter == "" {
		println("[WARN] colors: format: begin formatter is missing, Prepare() not called?")
	}
}

// Print print with fmt.Print() colored.
func (f Format) Print(args ...any) {
	f.veirfy()
	fmt.Print(f.parse(args...)...)
}

// Println print with fmt.Println() colored.
func (f Format) Println(args ...any) {
	f.veirfy()
	fmt.Println(f.parse(args...)...)
}

// Printf print with fmt.Printf() colored.
func (f Format) Printf(format string, args ...any) {
	f.veirfy()
	fmt.Printf(fmt.Sprint(f.parse(format)...), args...)
}
