package color

import "testing"

func TestNewColor(t *testing.T) {
	e := Color{24, 1.0, 1.0}
	a := *NewColor("ff6500")

	if e != a {
		t.Errorf("got %v expected %v", a, e)
	}
}

func TestHsvString(t *testing.T) {
	tt := map[string]string {
		"274, 49, 35": HsvString(NewColor("#462D59")),
		"240, 21, 35": HsvString(NewColor("#464659")),
		"0, 0, 27": HsvString(NewColor("#464646")),
		"0, 0, 0": HsvString(&Color{0, 0, 0}),
	}
	for e, a := range tt {
		if e != a {
			t.Errorf("got %v expected %v", a, e)
		}
	}
}

func TestRgbString(t *testing.T) {
	tt := map[string]string {
		"0.00, 0.00, 0.00": RgbString(NewColor("000000")),
		"1.00, 0.00, 0.00": RgbString(NewColor("ff0000")),
		"0.00, 1.00, 0.00": RgbString(NewColor("00ff00")),
		"0.00, 0.00, 1.00": RgbString(NewColor("0000ff")),
		"1.00, 0.40, 0.00": RgbString(NewColor("#ff6500")),
	}
	for e, a := range tt {
		if e != a {
			t.Errorf("got %s expected %s", a, e)
		}
	}
}

func TestRgbString255(t *testing.T) {
	tt := map[string]string {
		"0, 0, 0": RgbString255(NewColor("000000")),
		"255, 0, 0": RgbString255(NewColor("ff0000")),
		"0, 255, 0": RgbString255(NewColor("00ff00")),
		"0, 0, 255": RgbString255(NewColor("0000ff")),
		"204, 207, 51": RgbString255(NewColor("CCCF33")),
		"the other thing": RgbString(NewColor("CCCF33")),
		// "225, 101, 0": RgbString255(NewColor("#ff6500")),
	}
	for e, a := range tt {
		if e != a {
			t.Errorf("got %s expected %s", a, e)
		}
	}
}