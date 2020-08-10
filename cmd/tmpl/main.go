package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/xen0ne/riceutils/pkg/color"
)

var (
	pal  *color.Pallet
	from io.Reader
	to   io.Writer
	p    *string
)

func init() {
	p = flag.String("pallet", "", "")
	flag.StringVar(p, "p", "", "")
}

func main() {
	flag.Parse()

	switch len(flag.Args()) {
	case 0:
		from = os.Stdin
		to = os.Stdout
	case 1:
		var err error
		from, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		to = os.Stdout
	case 2:
		var err error
		from, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// open/create other file
		to, err = os.OpenFile(flag.Arg(1), os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Print("this is run")
		from = os.Stdin
		to = os.Stdout
	}

	var pal *color.Pallet
	if *p != "" {
		// pallet from file
		pal = color.PalletFromName(*p)
	} else {
		pal = color.PalletFromName("default")
	}

	err := color.ApplyPallet(from,
		pal, to)
	if err != nil {
		fmt.Println(err)
	}
}