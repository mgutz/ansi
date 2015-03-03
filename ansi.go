package ansi

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	black = iota
	red
	green
	yellow
	blue
	magenta
	cyan
	white

	normalIntensityFG = 30
	highIntensityFG   = 90
	normalIntensityBG = 40
	highIntensityBG   = 100

	start     = "\033["
	bold      = "1;"
	blink     = "5;"
	underline = "4;"
	inverse   = "7;"

	// Reset is the ANSI reset escape sequence
	Reset = "\033[0m"
)

var (
	plain = false
	// Colors maps common color names to their ANSI color code.
	Colors = map[string]int{
		"black":   black,
		"red":     red,
		"green":   green,
		"yellow":  yellow,
		"blue":    blue,
		"magenta": magenta,
		"cyan":    cyan,
		"white":   white,
	}
)

// ColorCode returns the ANSI color color code for style.
func ColorCode(style string) string {
	return colorCode(style).String()
}

// Gets the ANSI color code for a style.
func colorCode(style string) *bytes.Buffer {
	buf := bytes.NewBufferString("")
	if plain || style == "" {
		return buf
	}
	if style == "reset" {
		buf.WriteString(Reset)
		return buf
	}

	foregroundBackground := strings.Split(style, ":")
	foreground := strings.Split(foregroundBackground[0], "+")
	fg := Colors[foreground[0]]
	fgStyle := ""
	if len(foreground) > 1 {
		fgStyle = foreground[1]
	}

	bg, bgStyle := "", ""

	if len(foregroundBackground) > 1 {
		background := strings.Split(foregroundBackground[1], "+")
		bg = background[0]
		if len(background) > 1 {
			bgStyle = background[1]
		}
	}

	buf.WriteString(start)
	base := normalIntensityFG
	if len(fgStyle) > 0 {
		if strings.Contains(fgStyle, "b") {
			buf.WriteString(bold)
		}
		if strings.Contains(fgStyle, "B") {
			buf.WriteString(blink)
		}
		if strings.Contains(fgStyle, "u") {
			buf.WriteString(underline)
		}
		if strings.Contains(fgStyle, "i") {
			buf.WriteString(inverse)
		}
		if strings.Contains(fgStyle, "h") {
			base = highIntensityFG
		}
	}
	fmt.Fprintf(buf, "%d;", base+fg)

	base = normalIntensityBG
	if len(bg) > 0 {
		if strings.Contains(bgStyle, "h") {
			base = highIntensityBG
		}
		fmt.Fprintf(buf, "%d;", base+Colors[bg])
	}

	// remove last ";"
	buf.Truncate(buf.Len() - 1)
	buf.WriteRune('m')
	return buf
}

// Color colors a string based on the ANSI color code for style.
func Color(s, style string) string {
	if plain || len(style) < 1 {
		return s
	}
	buf := colorCode(style)
	buf.WriteString(s)
	buf.WriteString(Reset)
	return buf.String()
}

// ColorFunc creates a closureto avoid ANSI color code calculation.
func ColorFunc(style string) func(string) string {
	if style == "" {
		return func(s string) string {
			return s
		}
	}
	color := ColorCode(style)
	return func(s string) string {
		if plain || s == "" {
			return s
		}
		buf := bytes.NewBufferString(color)
		buf.WriteString(s)
		buf.WriteString(Reset)
		result := buf.String()
		return result
	}
}

// DisableColors disables ANSI color codes. On by default.
func DisableColors(disable bool) {
	plain = disable
}
