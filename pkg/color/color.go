// Package github.com/xen0ne/riceutils/color provides the Color and Pallet
// structs for describing and grouping colors
package color

import (
	"fmt"
	"math"
	"strconv"
)


// An Hue Saturation Brighness representation of a color
type Color struct {
	h int
	s, v float64
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
	return fmt.Sprintf(
		"%1.2f, %1.2f, %1.2f",
		r, g, b)
}

// RgbString creates a comma separated string of the
// red green and blue channels of the given color (from 0 to 255)
func  RgbString255(c *Color) string {
	r, g, b := Rgb255(c)
	return fmt.Sprintf("%d, %d, %d", r, g, b)
}

// Rgb returns red, green, and blue values of a Color [0-1]
func Rgb(col *Color) (float64, float64, float64) {
	h := float64(col.h)
	v := col.v 
	c := v * col.s
	x := c * (1 - math.Abs(math.Mod(h/60.0, 2.0) - 1.0))
	m := v - c

	var r, g, b float64

	switch {
	case h >= 0 && h < 60:
		r = c; g = x; b =0
	case h >= 60 && h < 120:
		r = x; g = c; b = 0
	case h >= 120 && h < 180:
		r = 0; g = c; b = x
	case h >= 180 && h < 240:
		r = 0; g = x; b = c
	case h >= 240 && h < 300:
		r = x; g = 0; b = c
	default:
		r = c; g = 0; b = x
	}
	return (r + m), (m+g), (b + m)
}

func Rgb255(c *Color) (int, int, int) {
	r, g, b := Rgb(c)
	return int(math.Ceil(r * 100)/100 * 255),
		int(math.Round(g * 1000)/1000 * 255),
		int(math.Round(b * 1000)/1000 * 255)
}




// Hsv returns the hue, sturation, and value of a given Color
func Hsv(c *Color) (int, int, int) {
	// formula from https://en.wikipedia.org/wiki/HSL_and_HSV#From_RGB
	return c.h, int(math.Round(c.s * 100)), int(math.Round(c.v * 100))
}

// HslString creates a comma separated string of
// the hue, saturation, and lightness values of the given color
func HsvString(c *Color) string {
	h, s, v := Hsv(c)
	return fmt.Sprintf("%d, %d, %d", h, s, v)
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

func hex2hsv(hex string) (int, float64, float64) {
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


	return int(math.Round(h)),
		math.Round(s * 100)/100, math.Round(v * 100)/100
}

// New takes a string hexadecimal color and creates
// a new Color and returns an error.
func NewColor(cstr string) *Color {

	h, s, v := hex2hsv(cstr)

	return &Color{h, s, v}
}

