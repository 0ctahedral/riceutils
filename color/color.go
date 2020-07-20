// Package github.com/xen0ne/riceutils/color provides the Color and Pallet
// structs for describing and grouping colors
package color

import (
	"fmt"
	"os"
	"strconv"
)

// the escape codes needed for terminal stuff
const (
	tesc = "\033Ptmux;\033\033]"
	tcesc ="\007\033\\"
	esc = "\033]"
	cesc = "\007"
)

// A 256 RGB representation of a color
type Color struct {
	R, G, B uint8
}

// Hex creates a string of the hexadecimal representation
// of the given Color
func Hex(c *Color) string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// Rgb creates a comma separated string of the red green and blue channels of the given color (from 0.0 to 1
func  Rgb(c *Color) string {
	r := float32(c.R ) / 255;
	g := float32(c.G ) / 255;
	b := float32(c.B ) / 255;
	return fmt.Sprintf("%1.2f, %1.2f, %1.2f", r, g, b)
}

// New takes a string hexadecimal color and creates
// a new Color and returns an error.
func NewColor(cstr string) *Color {
	c := ""
	i := 0
	clr := make([]uint8, 3)

	for _, l := range cstr {
		if l == '#' {
			continue
		}

		switch len(c) {
		case 0:
			c += string(l)
		case 1:
			c += string(l)
			num, err := strconv.ParseUint(c, 16, 8)
			if err != nil {
				panic(err)
			}
			clr[i] = uint8(num)
			i++
		case 2:
			c = string(l)
		}
	}
	return &Color{clr[0], clr[1], clr[2]}
}
/*
    [ -n "$TMUX" ] && {

        $1 == "bg" {
            printf "%s11;%s%s\n", esc, $2, cesc
            printf "%s4;0;%s%s\n", esc, $2, cesc
        }
        $1 == "bg_alt" {
            printf "%s;8;%s%s\n", esc, $2, cesc
        }
        $1 == "primay" {
            printf "%s4;1;%s%s\n", esc, $2, cesc
            printf "%s4;9;%s%s\n", esc, $2, cesc
        }
        $1 == "secondary" {
            printf "%s4;2;%s%s\n", esc, $2, cesc
            printf "%s4;10;%s%s\n", esc, $2, cesc
        }
        $1 == "alert" {
            printf "%s4;3;%s%s\n", esc, $2, cesc
            printf "%s4;11;%s%s\n", esc, $2, cesc
        }
        $1 == "cursor" {
            printf "%s4;4;%s%s\n", esc, $2, cesc
            printf "%s4;12;%s%s\n", esc, $2, cesc
            printf "%s12;%s%s\n", esc, $2, cesc
        }
        $1 == "fill" {
            printf "%s4;5;%s%s\n", esc, $2, cesc
            printf "%s4;6;%s%s\n", esc, $2, cesc
            printf "%s4;13;%s%s\n", esc, $2, cesc
            printf "%s4;14;%s%s\n", esc, $2, cesc
        }
        $1 == "fg" {
            printf "%s10;%s%s\n", esc, $2, cesc
            printf "%s4;7;%s%s\n",esc,  $2, cesc
        }
        $1 == "fg_alt" {
            printf "%s17;%s%s\n", esc, $2, cesc
            printf "%s4;15;%s%s\n", esc, $2, cesc
        }
*/

// Escape returns the full escape code to change a terminal's property given
// by p to the color given by c
func Escape(c *Color, p string) string {
	oe := esc
	ce := cesc
	// if tmux variable is set then we need extra escape codes
	if os.Getenv("TMUX") != "" {
		oe = tesc
		ce = tcesc
	}
    return fmt.Sprintf("%s%s%s%s\n", oe, p, Hex(c), ce)
}
