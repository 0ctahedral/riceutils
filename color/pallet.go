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
		bg:		NewColor("#ffffff"),
		bg_alt:	NewColor("#ffffff"),
		pri:	NewColor("#ffffff"),
		sec:	NewColor("#ffffff"),
		alert:	NewColor("#ffffff"),
		cur:	NewColor("#ffffff"),
		fill1:	NewColor("#ffffff"),
		fill2:	NewColor("#ffffff"),
		fg:		NewColor("#ffffff"),
		fg_alt:	NewColor("#ffffff"),
	}
}

// DefaultPallet fills a pallet with default values inspired by palenight
func DefaultPallet() *Pallet {
	return &Pallet{
		bg:		NewColor("#292D3E"),
		bg_alt:	NewColor("#697098"),
		pri:	NewColor("#c792ea"),
		sec:	NewColor("#C4E88D"),
		alert:	NewColor("#ff869a"),
		cur:	NewColor("#FFCB6B"),
		fill1:	NewColor("#82b1ff"),
		fill2:	NewColor("#939ede"),
		fg:		NewColor("#dde3eb"),
		fg_alt:	NewColor("#C7D8FF"),
	}
}
