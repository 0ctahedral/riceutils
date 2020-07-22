// Package github.com/xen0ne/riceutils/temp provides functions for applying Pallets to golang templates
package templ

import (
	"io/ioutil"
	"text/template"
	"io"
	"github.com/xen0ne/riceutils/pkg/color"
)

// ApplyPallet reads a template from a reader, applies the given pallet to it
// and then writes the filled in template to the writer.
func ApplyPallet(r io.Reader, p *color.Pallet, w io.Writer) error {

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	funcs := template.FuncMap{
		"hex": color.HexString,
		"rgb": color.RgbString,
		"rgb225": color.RgbString225,
		"hsv": color.HsvString,
	}

	tmpl := template.Must(template.New("test").Funcs(funcs).Parse(string(b)))

	err = tmpl.Execute(w, p.Iter())
	if err != nil {
		return err
	}

	return nil
}
