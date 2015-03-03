package main

import (
	"fmt"

	"github.com/mattn/go-colorable"
	"github.com/mgutz/ansi"
)

func main() {
	printColors()
}

func pad(s string, length int) string {
	for len(s) < length {
		s += " "
	}
	return s
}

func padColor(s string, styles []string) string {
	buffer := ""
	for _, style := range styles {
		buffer += ansi.Color(pad(s+style, 20), s+style)
	}
	return buffer
}

func printPlain() {
	ansi.DisableColors(true)
	bgColors := []string{
		"",
		":black",
		":red",
		":green",
		":yellow",
		":blue",
		":magenta",
		":cyan",
		":white",
	}
	for fg := range ansi.Colors {
		for _, bg := range bgColors {
			println(padColor(fg, []string{"" + bg, "+b" + bg, "+bh" + bg, "+u" + bg}))
			println(padColor(fg, []string{"+uh" + bg, "+B" + bg, "+Bb" + bg /* backgrounds */, "" + bg + "+h"}))
			println(padColor(fg, []string{"+b" + bg + "+h", "+bh" + bg + "+h", "+u" + bg + "+h", "+uh" + bg + "+h"}))
		}
	}
}

func printColors() {
	ansi.DisableColors(false)
	stdout := colorable.NewColorableStdout()

	bgColors := []string{
		"",
		":black",
		":red",
		":green",
		":yellow",
		":blue",
		":magenta",
		":cyan",
		":white",
	}
	for fg := range ansi.Colors {
		for _, bg := range bgColors {
			fmt.Fprintln(stdout, padColor(fg, []string{"" + bg, "+b" + bg, "+bh" + bg, "+u" + bg}))
			fmt.Fprintln(stdout, padColor(fg, []string{"+uh" + bg, "+B" + bg, "+Bb" + bg /* backgrounds */, "" + bg + "+h"}))
			fmt.Fprintln(stdout, padColor(fg, []string{"+b" + bg + "+h", "+bh" + bg + "+h", "+u" + bg + "+h", "+uh" + bg + "+h"}))
		}
	}
}
