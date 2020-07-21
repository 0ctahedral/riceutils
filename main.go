package main

import (
	// "os"
	// "strings"
	"fmt"
	"github.com/xen0ne/riceutils/color"
	// "github.com/xen0ne/riceutils/temp"
	"github.com/xen0ne/riceutils/term"
)



func main() {
	const t = `my fave color is {{hex .fg }}
my fave color in rbg {{rgb225 .fg }}
my fave color in rbg {{rgb .fg }}
my fave color in hsv {{hsv .fg }}`

	simp := color.CleanPallet()

	// err := temp.ApplyPallet(strings.NewReader(t),
	// 	simp, os.Stdout)
	// if err != nil { println(err) }

	fmt.Printf("%s", term.EscPallet(simp, term.Stdmap))
}
