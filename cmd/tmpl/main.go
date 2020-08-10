package main

import (
	"fmt"
	"io"
	"os"

	"github.com/xen0ne/riceutils/pkg/color"
)

var (
	pal  *color.Pallet
	from io.Reader
	to   io.Writer
)

func main() {
	// how many arguments are there?
	// 0 read from stdin
	// 1 read pallet from file name output to stdout
	// 2 read pallet from file and output to second file

	switch len(os.Args) {
	case 1:
		from = os.Stdin
		to = os.Stdout
	case 2:
		var err error
		from, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		to = os.Stdout
	default:
		from = os.Stdin
		to = os.Stdout
	}

	simp := color.CleanPallet()

	err := color.ApplyPallet(from,
		simp, to)
	if err != nil {
		fmt.Println(err)
	}
}
