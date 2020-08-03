package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/xen0ne/riceutils/pkg/color"
	"github.com/xen0ne/riceutils/pkg/term"
)

var args = make(map[string]*string)

func arginit() {
	args["x"] = flag.String("hex", "", "")
	flag.StringVar(args["x"], "x", "", "")
	args["r"] = flag.String("rgb", "", "")
	flag.StringVar(args["r"], "r", "", "")
	args["s"] = flag.String("hsv", "", "")
	flag.StringVar(args["s"], "s", "", "")

	flag.Usage = func() {
		use := `usage: pallet [options...] [pallet]

supplying no arguments will apply the current pallet to the terminal
supplying only a pallet will apply the given pallet to the terminal

general options:
-h | --help prints this message

printing options:
by using a -- to supply no arguments will print all colors from the pallet
-x | --hex prints the given color in hex format
-r | --rgb prints the given color in comma separated red, green, and blue
-s | --hsv prints the given color in comma separated hue, saturation, and value
	`
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", use)
	}
}

func main() {

	arginit()
	
	flag.Parse()

	if len(os.Args) == 1 {
		def := color.DefaultPallet()
		fmt.Printf("%s", term.EscPallet(def, term.Stdmap))
		os.Exit(0)
	}

	if len(os.Args) == 2 {
		p := PalletFromFile()
		fmt.Printf("%s", term.EscPallet(p, term.Stdmap))
		os.Exit(0)
	}

	if *args["x"] != "" {
		if (*args["x"] == "-" || *args["x"] == "--") {
			PrintPallet(color.HexString)
			os.Exit(0)
		} else {
			PrintColorFromPallet(*args["x"], color.HexString)
		}
	}

	if *args["r"] != "" {
		if (*args["r"] == "-" || *args["r"] == "--") {
			PrintPallet(color.RgbString255)
			os.Exit(0)
		} else {
			PrintColorFromPallet(*args["r"], color.RgbString255)
		}
	}

	if *args["s"] != "" {
		if (*args["s"] == "-" || *args["s"] == "--") {
			PrintPallet(color.HsvString)
			os.Exit(0)
		} else {
			PrintColorFromPallet(*args["s"], color.HsvString)
		}
	}
}

func PalletFromFile() *color.Pallet {
	path := os.Getenv("PALLET_PATH")
	fp := os.ExpandEnv(fmt.Sprintf("%s%s", path, flag.Args()[0]))
	f, err := os.Open(fp)
	if err != nil {
		//fmt.Printf("file %s does not exist")
		fmt.Println(err)
		os.Exit(1)
	}
	p, err := color.ParseReader(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return p
}

func PrintPallet(a func(*color.Color) string) {
	var p *color.Pallet

	if len(flag.Args()) == 0 {
		p = color.DefaultPallet()
	} else {
		// read the pallet from file
		p = PalletFromFile()
	}


	for _, v := range p.Iter() {
		fmt.Println(a(v))
	}

}

func PrintColorFromPallet(str string, a func(*color.Color) string) {
	var p *color.Pallet

	if len(flag.Args()) == 0 {
		p = color.DefaultPallet()
	} else {
		p = PalletFromFile()
	}

	if c, ok := p.Iter()[str]; ok {
		fmt.Println(a(c))
		os.Exit(0)
	}
	fmt.Fprintf(os.Stderr, "%s is not a valid color\n", str)
	os.Exit(1)
}