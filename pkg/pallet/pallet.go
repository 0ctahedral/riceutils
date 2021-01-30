package pallet

import (
	"bufio"
	"errors"
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
	bg,
	bg_alt,
	fg,
	fg_alt,
	pri,
	sec,
	alert,
	cur,
	com,
	block *color.Color
}

// Iter returns a copy of the Pallet in map form for iterating
func (p *Pallet) Iter() map[string]*color.Color {
	return map[string]*color.Color{
		"bg":        p.bg,
		"bg_alt":    p.bg_alt,
		"fg":        p.fg,
		"fg_alt":    p.fg_alt,
		"pri":       p.pri,
		"sec":       p.sec,
		"primary":   p.pri,
		"secondary": p.sec,
		"alert":     p.alert,
		"cur":       p.cur,
		"cursor":    p.cur,
		"com":     p.com,
		"block":     p.block,
	}
}

func (p *Pallet) ChangeColor(str string, c *color.Color) error {
	switch str {
	case "bg":
		p.bg = c
	case "bg_alt":
		p.bg_alt = c
	case "fg":
		p.fg = c
	case "fg_alt":
		p.fg_alt = c
	case "pri", "primary":
		p.pri = c
	case "sec", "secondary":
		p.sec = c
	case "alert":
		p.alert = c
	case "cur", "cursor":
		p.cur = c
	case "com":
		p.com = c
	case "block":
		p.block = c
	default:
		return errors.New(fmt.Sprintf("Invalid color name: %s", str))
	}

	return nil
}

// CleanPallet fills a pallet with white values
func CleanPallet() *Pallet {
	return &Pallet{
		bg:     color.NewColor("#ffffff"),
		bg_alt: color.NewColor("#ffffff"),
		pri:    color.NewColor("#ffffff"),
		sec:    color.NewColor("#ffffff"),
		alert:  color.NewColor("#ffffff"),
		cur:    color.NewColor("#ffffff"),
		com:  color.NewColor("#ffffff"),
		block:  color.NewColor("#ffffff"),
		fg:     color.NewColor("#ffffff"),
		fg_alt: color.NewColor("#ffffff"),
	}
}

// DefaultPallet fills a pallet with default values inspired by palenight
func DefaultPallet() *Pallet {
	return &Pallet{
		bg:     color.NewColor("#292D3E"),
		bg_alt: color.NewColor("#697098"),
		pri:    color.NewColor("#c792ea"),
		sec:    color.NewColor("#C4E88D"),
		alert:  color.NewColor("#ff869a"),
		cur:    color.NewColor("#FFCB6B"),
		com:  color.NewColor("#82b1ff"),
		block:  color.NewColor("#939ede"),
		fg:     color.NewColor("#dde3eb"),
		fg_alt: color.NewColor("#C7D8FF"),
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
		os.Exit(1)
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

	tmpl := template.Must(template.New("test").Funcs(funcs).Parse(string(b)))

	err = tmpl.Execute(w, p.Iter())
	if err != nil {
		return err
	}

	return nil
}
