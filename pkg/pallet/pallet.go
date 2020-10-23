package pallet

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/xen0ne/riceutils/pkg/color"
)

// A Pallet is a map of string color names to color.Colors
type Pallet struct {
	cmap map[string]*color.Color
}

// Map returns a copy of the Pallet in map form for iterating
func (p *Pallet) Map() map[string]*color.Color {
	return p.cmap
}

// ChangeColor changes a color if it exists and returns an error otherwise
func (p *Pallet) ChangeColor(str string, c *color.Color) error {
	// check if the color exists
	if _, ok := p.cmap[str]; ok {

	} else {
		return fmt.Errorf("Invalid color name: %s", str)
	}

	return nil
}

// CleanPallet fills a pallet with white values
func CleanPallet() *Pallet {
	return &Pallet{
		cmap: map[string]*color.Color{
			"bg":    color.NewColor("#ffffff"),
			"bg+":   color.NewColor("#ffffff"),
			"pri":   color.NewColor("#ffffff"),
			"sec":   color.NewColor("#ffffff"),
			"alert": color.NewColor("#ffffff"),
			"cur":   color.NewColor("#ffffff"),
			"fill1": color.NewColor("#ffffff"),
			"fill2": color.NewColor("#ffffff"),
			"fg":    color.NewColor("#ffffff"),
			"fg+":   color.NewColor("#ffffff"),
		},
	}
}

// DefaultPallet fills a pallet with default values inspired by palenight
func DefaultPallet() *Pallet {
	return &Pallet{
		cmap: map[string]*color.Color{
			"bg":    color.NewColor("#292D3E"),
			"bg+":   color.NewColor("#697098"),
			"pri":   color.NewColor("#c792ea"),
			"sec":   color.NewColor("#C4E88D"),
			"alert": color.NewColor("#ff869a"),
			"cur":   color.NewColor("#FFCB6B"),
			"fill1": color.NewColor("#82b1ff"),
			"fill2": color.NewColor("#939ede"),
			"fg":    color.NewColor("#dde3eb"),
			"fg+":   color.NewColor("#C7D8FF"),
		},
	}
}

func ParseReader(r io.Reader) (*Pallet, error) {
	// test if regex works
	reg := regexp.MustCompile(`^(\w+)[[:space:]]*:.*\"*(#\w+)\"*$`)
	s := bufio.NewScanner(r)

	p := CleanPallet()
	for s.Scan() {
		m := reg.FindStringSubmatch(s.Text())
		if len(m) == 3 {
			if err := p.ChangeColor(m[1], color.NewColor(m[2])); err != nil {
				fmt.Println(err)
			}
		}
	}

	return p, nil
}

// PalletFromName reads a a pallet from a file in the
// PALLET_PATH with the same name as pname
func PalletFromName(pname string) *Pallet {
	path := PalletPath()
	fp := os.ExpandEnv(fmt.Sprintf("%s/%s", path, pname))
	f, err := os.Open(fp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot find pallet %s\n%s\n",
			pname, "hint: is it in your PALLT_PATH?")
		// os.Exit(1)
		fmt.Fprintf(os.Stderr, "resorting to default pallet\n")
		return DefaultPallet()
	}

	p, err := ParseReader(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return p
}

// Resolves the pallet path or returns the default
func PalletPath() string {
	// TODO: determine if relative path
	path := os.Getenv("PALLET_PATH")
	if path == "" {
		path = "$HOME/.config/pallets/"
	}
	return strings.TrimRight(path, "/")
}

// ApplyPallet reads a template from a reader, applies the given pallet to it
// and then writes the filled in template to the writer.
func ApplyPallet(r io.Reader, p *Pallet, w io.Writer) error {

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	funcs := template.FuncMap{
		"hex":    color.HexString,
		"rgb":    color.RgbString,
		"rgb225": color.RgbString255,
		"hsv":    color.HsvString,
	}

	tmpl, err := template.New("test").Funcs(funcs).Parse(string(b))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	err = tmpl.Execute(w, p.Map())
	if err != nil {
		return err
	}

	return nil
}
