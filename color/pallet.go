package color
import (
	"text/template"
	"io/ioutil"
	"io"
)

type Pallet struct {
	 Primary Color
	 Secondary Color
}

// ApplyPallet reads a template from a reader, applies the given pallet to it
// and then writes the filled in template to the writer
func ApplyPallet(r io.Reader, p Pallet, w io.Writer) error {

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	funcs := template.FuncMap{
		"hex": ToHex,
		"rgb": ToRgb,
	}

	tmpl := template.Must(template.New("test").Funcs(funcs).Parse(string(b)))

	err = tmpl.Execute(w, p)
	if err != nil {
		return err
	}

	return nil
}
