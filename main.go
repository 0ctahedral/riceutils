package main

import (
	"os"
	"strings"
	"github.com/xen0ne/ricemgr/color"
)



func main() {
	const temp = "my fave color is {{hex .Primary }}\nmy second fave color is {{hex .Secondary }}\n"

	simp := color.Pallet{color.Color{0, 0xff, 0}, color.Color{0, 0, 0xff}}

	err := color.ApplyPallet(strings.NewReader(temp), simp, os.Stdout)
	if err != nil {
		panic(err)
	}
}
