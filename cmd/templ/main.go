package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/xen0ne/riceutils/color"
	"github.com/xen0ne/riceutils/templ"
)

// cmdline args
func usage() {
	use = `apply <pallet> -i -o
	`
	fmt.Println(use)
}

func main() {
	const t = `my fave color is {{hex .fg }}
	my fave color in rbg {{rgb225 .fg }}
	my fave color in rbg {{rgb .fg }}
	my fave color in hsv {{hsv .fg }}`

	simp := color.CleanPallet()

	err := templ.ApplyPallet(strings.NewReader(t),
		simp, os.Stdout)
	if err != nil { fmt.Println(err) }
}
