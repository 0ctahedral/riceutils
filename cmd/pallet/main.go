package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/xen0ne/riceutils/pkg/color"
	"github.com/xen0ne/riceutils/pkg/term"
)


// cmdline args
func usage() {
	use := `pallet <pallet>
supplying no arguments will apply the current pallet to the terminal
-h | --help prints this message
-n | --new creates a new pallet with the given name
-e | --edit changes a value in the pallet (idk how to do that yet)
-p | --picker opens a gui for editing the given pallet
	`
	fmt.Println(use)
}

func main() {

	//var args map[string]*bool
	
	// n := flag.Bool("new", false, "name of new pallet")
	h := flag.Bool("help", false, "help menu")
	flag.Parse()

	if len(os.Args) == 1 {
		simp := color.DefaultPallet()
		fmt.Printf("%s", term.EscPallet(simp, term.Stdmap))
	}

	if (*h) {
		usage()
	}
}
