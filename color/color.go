// Package github.com/xen0ne/riceutils/color provides the Color and Pallet
// structs for describing and grouping colors
package color

import (
	"fmt"
	"os"
	"math"
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

// HexString creates a string of the hexadecimal representation
// of the given Color
func HexString(c *Color) string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// RgbString creates a comma separated string of the
// red green and blue channels of the given color (from 0.0 to 1
func  RgbString(c *Color) string {
	r, g, b := Rgb(c)
	return fmt.Sprintf("%1.2f, %1.2f, %1.2f", r, g, b)
}

// RgbString creates a comma separated string of the
// red green and blue channels of the given color (from 0 to 225)
func  RgbString225(c *Color) string {
	return fmt.Sprintf("%d, %d, %d", c.R, c.G, c.B)
}

// Rgb returns r, g, and b values of a Color [0-1]
func Rgb(c *Color) (r,g,b float32) {
	r = float32(c.R ) / 255;
	g = float32(c.G ) / 255;
	b = float32(c.B ) / 255;
	return
}

// HslString creates a comma separated string of
// the hue, saturation, and lightness values of the given color
func HsvString(color *Color) string {
	// formula from https://en.wikipedia.org/wiki/HSL_and_HSV#From_RGB
	r, g, b := Rgb(color)
	var max float32 = 0
	var min float32 = 1
	for _, f := range []float32{r, g, b} {
		if (f > max ) {max = f}
		if (f < min ) {min = f}
	}


	c := max - min
	v := max
	var s float32
	if v == 0 {
		s = 0
	} else {
		s = c/v
	}

	var h float32
	if c == 0 {
		h = 0
	} else {
		switch v {
		case r:
			h = 60 * (g-b)/c
		case g:
			h = 60 * (2 + (b-r)/c)
		case b:
			h = 60 * (4 + (r-g)/c)
		default:
			fmt.Printf("\nmax: %f\n", max)
			fmt.Printf("min: %f\n", min)
			fmt.Printf("%f, %f, %f \n", r, g, b)
		}
	}

	//if h < 0 { h = 360 - h }

	s = float32(math.Round(float64(s * 100)))
	v = float32(math.Round(float64(v * 100)))

	return fmt.Sprintf("%d, %d, %d",
		int(h), int(s), int(v))
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
    return fmt.Sprintf("%s%s%s%s\n", oe, p, HexString(c), ce)
}
