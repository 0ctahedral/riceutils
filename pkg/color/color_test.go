package color

import "testing"

func TestHsvString(t *testing.T) {
	tt := map[string]string {
		"274, 49, 35": HsvString(NewColor("#462D59")),
		"240, 21, 35": HsvString(NewColor("#464659")),
		"0, 0, 27": HsvString(NewColor("#464646")),
		"0, 0, 0": HsvString(&Color{0, 0, 0}),
	}
	for e, a := range tt {
		if e != a {
			t.Errorf("go %s expected %s", a, e)
		}
	}
}

func TestRgbString(t *testing.T) {
	tt := map[string]string {
		"0.00, 0.00, 0.00": RgbString(&Color{0, 0, 0}),
		"1.00, 0.00, 0.00": RgbString(&Color{255, 0, 0}),
		"0.00, 1.00, 0.00": RgbString(&Color{0, 255, 0}),
		"0.00, 0.00, 1.00": RgbString(&Color{0, 0, 255}),
		"1.00, 0.40, 0.00": RgbString(NewColor("#ff6500")),
	}
	for e, a := range tt {
		if e != a {
			t.Errorf("go %s expected %s", a, e)
		}
	}
}

func TestNewColor(t *testing.T) {
	tt := map[Color]Color {
		Color{0, 255, 0}: *NewColor("#00ff00"),
		Color{0, 255, 0}: *NewColor("00ff00"),
		Color{0, 255, 0}: *NewColor("00Ff00"),
		Color{0, 255, 0}: *NewColor("00FF00"),
		Color{0, 255, 0}: *NewColor("#00ff00"),
		Color{0, 255, 0}: *NewColor("#00Ff00"),
		Color{0, 255, 0}: *NewColor("#00FF00"),
		Color{0, 255, 0}: *NewColor("#00FF00ff"),
	}
	for e, a := range tt {
		if e != a {
			t.Errorf("go %v expected %v", a, e)
		}
	}
}
