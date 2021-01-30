package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/xen0ne/riceutils/pkg/pallet"
)

var (
	pal  *pallet.Pallet
	from io.Reader
	// to   io.Writer
	p *string
)

func init() {
	p = flag.String("pallet", "", "")
	flag.StringVar(p, "p", "", "")
}

func main() {
	flag.Parse()

	// check for stdin
	inStat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open stdin because: %s\n", err.Error())
		os.Exit(1)
	}

	// check if there is data actually piped in
	if inStat.Mode()&os.ModeCharDevice == 0 {
		from = os.Stdin
	} else {
		// check for argument
		if len(flag.Args()) == 0 {
			fmt.Fprintln(os.Stderr, "no template file specified")
			os.Exit(1)
		}
		// open the first file for now
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
		from = io.Reader(f)
	}

	var pal *pallet.Pallet
	if *p != "" {
		// pallet from file
		pal = pallet.PalletFromName(*p)
	} else {
		pal = pallet.PalletFromName("default")
	}

	err = pallet.ApplyPallet(from,
		pal, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
