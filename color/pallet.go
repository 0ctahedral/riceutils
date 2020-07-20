package color

import (
)

// A Pallet is a map of string color names to Colors
type Pallet map[string]*Color

// CleanPallet fills a pallet with white values
func CleanPallet() Pallet {
	return Pallet{
		"bg":		&Color{255, 255, 255},
		"bg_alt":	&Color{255, 255, 255},
		"fg":		&Color{255, 255, 255},
		"fg_alt":	&Color{255, 255, 255},
		"pri":		&Color{255, 255, 255},
		"sec":		&Color{255, 255, 255},
		"alet":		&Color{255, 255, 255},
		"cur":		&Color{255, 255, 255},
		"fill1":	&Color{255, 255, 255},
		"fill2":	&Color{255, 255, 255},
		"fill3":	&Color{255, 255, 255},
	}
}
