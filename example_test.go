package ansi_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mgutz/ansi"
)

func Example() {
	// colorize a string, SLOW
	msg := ansi.Color("foo", "red+b:white")

	// create a FAST closure function to avoid computation of ANSI code
	phosphorize := ansi.ColorFunc("green+h:black")
	msg = phosphorize("Bring back the 80s!")
	msg2 := phosphorize("Look, I'm a CRT!")

	// cache escape codes and build strings manually
	lime := ansi.ColorCode("green+h:black")
	reset := ansi.ColorCode("reset")

	fmt.Println(lime, "Bring back the 80s!", reset)

	_ = msg
	_ = msg2
}

func ExampleColor() {
	s := "example"

	_ = ansi.Color(s, "red")
	_ = ansi.Color(s, "red+d")
	_ = ansi.Color(s, "red+b")
	_ = ansi.Color(s, "red+B")
	_ = ansi.Color(s, "red+u")
	_ = ansi.Color(s, "red+bh")
	_ = ansi.Color(s, "red:white")
	_ = ansi.Color(s, "red+b:white+h")
	_ = ansi.Color(s, "red+B:white+h")
	_ = ansi.Color(s, "off")
}

func TestReadmeExamples(t *testing.T) {
	ansi.DisableColors(false)

	// Example 1: basic Color
	msg := ansi.Color("foo", "red+b:white")
	if !strings.Contains(msg, "foo") {
		t.Errorf("Color failed to include text: %q", msg)
	}

	// Example 2: ColorFunc closure
	phosphorize := ansi.ColorFunc("green+h:black")
	msg = phosphorize("Bring back the 80s!")
	if !strings.Contains(msg, "Bring back the 80s!") {
		t.Errorf("phosphorize failed to include text: %q", msg)
	}
	msg2 := phosphorize("Look, I'm a CRT!")
	if !strings.Contains(msg2, "Look, I'm a CRT!") {
		t.Errorf("phosphorize failed to include text: %q", msg2)
	}

	// Example 3: ColorCode manual composition
	lime := ansi.ColorCode("green+h:black")
	reset := ansi.ColorCode("reset")
	if lime == "" || reset == "" {
		t.Errorf("expected non-empty escape codes, got lime=%q, reset=%q", lime, reset)
	}

	composed := fmt.Sprint(lime, "Bring back the 80s!", reset)
	if !strings.Contains(composed, "Bring back the 80s!") {
		t.Errorf("composed string missing text: %q", composed)
	}

	// Other examples from README
	s := "test string"
	styles := []string{
		"red",
		"red+d",
		"red+b",
		"red+B",
		"red+u",
		"red+bh",
		"red:white",
		"red+b:white+h",
		"red+B:white+h",
		"off",
	}

	for _, style := range styles {
		result := ansi.Color(s, style)
		if style == "off" {
			expected := s + ansi.Reset
			if result != expected {
				t.Errorf("Color(s, 'off') = %q, want %q", result, expected)
			}
		} else {
			if !strings.Contains(result, s) {
				t.Errorf("Color(s, %q) missing raw string: %q", style, result)
			}
		}
	}
}
