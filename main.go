package main

import (
	"os"
	"strings"
	"github.com/xen0ne/riceutils/color"
	"github.com/xen0ne/riceutils/temp"
)



func main() {
	const t = "my fave color is {{hex .bg }}\nmy second fave color is {{rgb .fg }}\n"

	//simp := color.CleanPallet()

	err := temp.ApplyPallet(strings.NewReader(t), color.CleanPallet(), os.Stdout)
	if err != nil { println(err) }

	//fmt.Printf("%s", color.Escape(simp.Primary, "11;"))
}
