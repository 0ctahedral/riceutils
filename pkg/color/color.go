// Package github.com/xen0ne/riceutils/color provides the Color and Pallet
// structs for describing and grouping colors
package color

import (
	"fmt"
	"math"
	"strconv"
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
func  RgbString255(c *Color) string {
	return fmt.Sprintf("%d, %d, %d", c.R, c.G, c.B)
}

// Rgb returns red, green, and blue values of a Color [0-1]
func Rgb(c *Color) (r,g,b float32) {
	r = float32(c.R ) / 255;
	g = float32(c.G ) / 255;
	b = float32(c.B ) / 255;
	return
}

// Hsv returns the hue, sturation, and value of a given Color
func Hsv(color *Color) (int, int, int) {
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

	s = float32(math.Round(float64(s * 100)))
	v = float32(math.Round(float64(v * 100)))

	return int(h), int(s), int(v)
}

// HslString creates a comma separated string of
// the hue, saturation, and lightness values of the given color
func HsvString(c *Color) string {
	h, s, v := Hsv(c)
	return fmt.Sprintf("%d, %d, %d", h, s, v)
}

// New takes a string hexadecimal color and creates
// a new Color and returns an error.
func NewColor(cstr string) *Color {
	c := ""
	i := 0
	clr := make([]uint8, 3)

	for _, l := range cstr {
		// we can disregard alpha
		if i == 3 { break }

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

