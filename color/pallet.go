package color

import (
)

// A Pallet is a map of string color names to Colors
type Pallet struct {
		bg,
		bg_alt,
		fg,
		fg_alt,
		pri,
		sec,
		alert,
		cur,
		fill1,
		fill2 *Color
}

// Iter returns a copy of the Pallet in map form for iterating
func (p *Pallet) Iter() map[string]*Color {
	return map[string]*Color {
		"bg":		p.bg,
		"bg_alt":	p.bg_alt,
		"fg":		p.fg,
		"fg_alt":	p.fg_alt,
		"pri":		p.pri,
		"sec":		p.sec,
		"alert":	p.alert,
		"cur":		p.cur,
		"fill1":	p.fill1,
		"fill2":	p.fill2,
	}
}

// CleanPallet fills a pallet with white values
func CleanPallet() *Pallet {
	return &Pallet{
		bg:		&Color{255, 255, 255},
		bg_alt:	&Color{255, 255, 255},
		pri:	&Color{255, 255, 255},
		sec:	&Color{255, 255, 255},
		alert:	&Color{255, 255, 255},
		cur:	&Color{255, 255, 255},
		fill1:	&Color{255, 255, 255},
		fill2:	&Color{255, 255, 255},
		fg:		NewColor("#00ff00"),
		fg_alt:	&Color{255, 255, 255},
	}
}
