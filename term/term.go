// Package github.com/xen0ne/riceutils/term provides utilies
// for printing Colors and Pallets to the screen
package term

import (
	"github.com/xen0ne/riceutils/color"
	"fmt"
	"os"
)

// the escape codes needed for terminal stuff
const (
	tesc = "\033Ptmux;\033\033]"
	tcesc ="\007\033\\"
	esc = "\033]"
	cesc = "\007"
)

var Stdmap = map[string][]string {
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

/*
    [ -n "$TMUX" ] && {

        $1 == "bg" {
            printf "%s11;%s%s\n", esc, $2, cesc
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
func EscPallet(p *color.Pallet, m map[string][]string) string {
	var ret string
	for k, c := range p.Iter() {
		fmt.Println(k)
		fmt.Println(color.HexString(c))
		if escs, ok := m[k]; ok {
			for _, e := range escs {
				ret += EscColor(c, e)
			}
		}
	}

	return ret
}

// PalletBlock prints a nice pallet for ya
func PalletBlock(p *color.Pallet) string {
	return ""
}
