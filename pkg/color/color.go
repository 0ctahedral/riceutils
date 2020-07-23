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
	h, s, v int
}

// HexString creates a string of the hexadecimal representation
// of the given Color
func HexString(c *Color) string {
	r, g, b := Rgb(c)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

// RgbString creates a comma separated string of the
// red green and blue channels of the given color (from 0.0 to 1
func  RgbString(c *Color) string {
	r, g, b := Rgb(c)
	return fmt.Sprintf("%1.2f, %1.2f, %1.2f",
		float32(r)/255, float32(g)/225, float32(b)/255)
}

// RgbString creates a comma separated string of the
// red green and blue channels of the given color (from 0 to 225)
func  RgbString225(c *Color) string {
	r, g, b := Rgb(c)
	return fmt.Sprintf("%d, %d, %d", r, g, b)
}

// Rgb returns red, green, and blue values of a Color [0-1]
func Rgb(col *Color) (int, int, int) {
	return	0, 0, 0
	/*
	var hp float64 = float64(col.h / 60.0)
	c := float64((col.v/100.0) * (col.s/100.0))
	x := c * (1 - math.Abs(math.Mod(hp,2.0) - 1.0))

	switch {
	case hp >= 0 && hp <= 1:
		return 1, 0, 0
	}
	return
	*/
}

// Hsv returns the hue, sturation, and value of a given Color
func Hsv(c *Color) (int, int, int) {
	// formula from https://en.wikipedia.org/wiki/HSL_and_HSV#From_RGB
	return c.h, c.s, c.v
}

// HslString creates a comma separated string of
// the hue, saturation, and lightness values of the given color
func HsvString(c *Color) string {
	return fmt.Sprintf("%d, %d, %d", c.h, c.s, c.v)
}

func hex2rgb(hex string) (int, int, int) {
	c := ""
	i := 0
	ret := make([]int, 3)

	for _, l := range hex {
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
			num, err := strconv.ParseInt(c, 16, 0)
			if err != nil {
				panic(err)
			}
			ret[i] = int(num)
			i++
		case 2:
			c = string(l)
		}
	}

	return ret[0], ret[1], ret[2]
}

func hex2hsv(hex string) (int, int, int) {
	ir, ig, ib := hex2rgb(hex)
	r := float64(ir)/255
	g := float64(ig)/255
	b := float64(ib)/255

	max := 0.0
	min := 1.0

	for _, f := range []float64{r, g, b} {
		if (f > max ) {max = f}
		if (f < min ) {min = f}
	}


	c := max - min
	v := max
	var s float64
	if v == 0 {
		s = 0
	} else {
		s = c/v
	}

	var h float64
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

	s = math.Round(float64(s * 100))
	v = math.Round(float64(v * 100))

	return int(math.Round(h)),
		int(math.Round(s)),
		int(math.Round(v))
}

// New takes a string hexadecimal color and creates
// a new Color and returns an error.
func NewColor(cstr string) *Color {

	h, s, v := hex2hsv(cstr)

	return &Color{h, s, v}
}

