package ansi

import (
	"fmt"
	"testing"
)

func TestPlain(t *testing.T) {
	DisableColors(true)
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
	for fg := range Colors {
		for _, bg := range bgColors {
			println(padColor(fg, []string{"" + bg, "+b" + bg, "+bh" + bg, "+u" + bg}))
			println(padColor(fg, []string{"+uh" + bg, "+B" + bg, "+Bb" + bg /* backgrounds */, "" + bg + "+h"}))
			println(padColor(fg, []string{"+b" + bg + "+h", "+bh" + bg + "+h", "+u" + bg + "+h", "+uh" + bg + "+h"}))
		}
	}
}

func TestStyles(t *testing.T) {
	PrintStyles()
	DisableColors(false)
	buf := colorCode("off")
	if buf.String() != "" {
		t.Fail()
	}
}

func TestFaint(t *testing.T) {
	actual := Color("foo", "white+f")
	expected := "\x1b[2;37mfoo\x1b[0m"
	if actual != "foo" {
		t.Fatalf("expected %#v, got %#v", expected, actual)
	}
}

func ExampleColorFunc() {
	brightGreen := ColorFunc("green+h")
	fmt.Println(brightGreen("lime"))
}
