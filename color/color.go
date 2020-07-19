package color

import (
	"fmt"
)

type Color struct {
	R, G, B uint8
}

// ToHex creates a string of the hexadecimal representation
// of the given Color
func ToHex(c Color) string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

func  ToRgb(c Color) string {
	return fmt.Sprintf("%d, %d, %d", c.R, c.G, c.B)
}
