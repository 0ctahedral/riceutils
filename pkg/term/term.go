// Package github.com/xen0ne/riceutils/term provides utilies
// for printing Colors and Pallets to the screen
package term

import (
	"fmt"
	"os"

	"github.com/xen0ne/riceutils/pkg/color"
	"github.com/xen0ne/riceutils/pkg/pallet"
)

// the escape codes needed for terminal stuff
const (
	tesc  = "\033Ptmux;\033\033]"
	tcesc = "\007\033\\"
	esc   = "\033]"
	cesc  = "\007"
)

var Stdmap = map[string][]string{
	"bg":     {"11;", "4;0;"},
	"bg_alt": {"4;8;"},
	"pri":    {"4;1;", "4;9;"},
	"sec":    {"4;2;", "4;10;"},
	"alert":  {"4;3;", "4;11;"},
	"cur":    {"4;4;", "4;12;", "12;"},
	"fill1":  {"4;5;", "4;13;"},
	"fill2":  {"4;6;", "4;14;"},
	"fg":     {"10;", "4;7;"},
	"fg_alt": {"17;", "4;15;"},
}

/*
var Ttymap = map[string][]string {
	"bg": {"11;", "4;0;"},
	"bg_alt": {"4;8;"},
	"pri": {"4;1;", "4;9;"},
	"sec": {"4;2;", "4;10;"},
	"alert": {"4;3;", "4;11;"},
	"cur": {"4;4;", "4;12;", "12;"},
	"fill1": {"4;5;", "4;13;"},
	"fill2": {"4;6;", "4;14;"},
	"fg": {"10;", "4;7;"},
	"fg_alt": {"17;", "4;15;"},
}
*/

// EscColor returns the full escape code to change a terminal's property given
// by p to the color given by c
func EscColor(c *color.Color, p string) string {
	oe := esc
	ce := cesc
	// if tmux variable is set then we need extra escape codes
	if os.Getenv("TMUX") != "" {
		oe = tesc
		ce = tcesc
	}
	return fmt.Sprintf("%s%s%s%s\n", oe, p, color.HexString(c), ce)
}

// EscPallet returns the escape codes for all the colors in a Pallet p
// based upon those provided by the map m
func EscPallet(p *pallet.Pallet, m map[string][]string) string {
	var ret string
	for k, c := range p.Map() {
		if escs, ok := m[k]; ok {
			for _, e := range escs {
				ret += EscColor(c, e)
			}
		}
	}

	ret += "\033[H"
	return ret
}

// PalletBlock prints a nice pallet for ya
func PalletBlock(p *pallet.Pallet) string {
	return ""
}
