package ansi

import (
	"fmt"
	"strings"
	"testing"
)

func TestPlain(t *testing.T) {
	DisableColors(true)
	PrintStyles()
}

func TestStyles(t *testing.T) {
	DisableColors(false)
	PrintStyles()
}

func TestDisableColors(t *testing.T) {
	fn := ColorFunc("red")

	buf := colorCode("off")
	if buf.String() != "" {
		t.Fail()
	}

	DisableColors(true)
	if Black != "" {
		t.Fail()
	}
	code := ColorCode("red")
	if code != "" {
		t.Fail()
	}
	s := fn("foo")
	if s != "foo" {
		t.Fail()
	}

	DisableColors(false)
	if Black == "" {
		t.Fail()
	}
	code = ColorCode("red")
	if code == "" {
		t.Fail()
	}
	// will have escape codes around it
	index := strings.Index(fn("foo"), "foo")
	if index <= 0 {
		t.Fail()
	}
}

func TestAttributeReset(t *testing.T) {
	boldRed := ColorCode("red+b")
	greenUnderline := ColorCode("green+u")
	s := fmt.Sprintf("normal %s bold red %s green underline %s", boldRed, greenUnderline, Reset)
	// See the results on the terminal for regression tests.
	fmt.Printf("Colored string: %s\n", s)
	fmt.Printf("Escaped string: %q\n", s)
	if s != "normal \x1b[0;1;31m bold red \x1b[0;4;32m green underline \x1b[0m" {
		t.Error("Attributes are not being reset")
	}
}

func TestHyperlink(t *testing.T) {
	defer DisableColors(false)
	DisableColors(false)

	tests := []struct {
		name     string
		url      string
		text     string
		expected string
	}{
		{
			name:     "standard URL and text",
			url:      "https://github.com",
			text:     "GitHub",
			expected: "\x1b]8;;https://github.com\x1b\\GitHub\x1b]8;;\x1b\\",
		},
		{
			name:     "URL with query parameters and fragment",
			url:      "https://example.com/path?foo=bar&baz=1#section",
			text:     "Deep Link",
			expected: "\x1b]8;;https://example.com/path?foo=bar&baz=1#section\x1b\\Deep Link\x1b]8;;\x1b\\",
		},
		{
			name:     "empty text",
			url:      "https://example.com",
			text:     "",
			expected: "\x1b]8;;https://example.com\x1b\\\x1b]8;;\x1b\\",
		},
		{
			name:     "empty URL",
			url:      "",
			text:     "text without url",
			expected: "\x1b]8;;\x1b\\text without url\x1b]8;;\x1b\\",
		},
		{
			name:     "both empty",
			url:      "",
			text:     "",
			expected: "\x1b]8;;\x1b\\\x1b]8;;\x1b\\",
		},
		{
			name:     "unicode URL and text",
			url:      "https://example.com/ümlaut",
			text:     "世界",
			expected: "\x1b]8;;https://example.com/ümlaut\x1b\\世界\x1b]8;;\x1b\\",
		},
		{
			name:     "colored text inside hyperlink",
			url:      "https://example.com",
			text:     Color("click here", "blue+u"),
			expected: "\x1b]8;;https://example.com\x1b\\" + Color("click here", "blue+u") + "\x1b]8;;\x1b\\",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual := Hyperlink(tc.url, tc.text)
			if actual != tc.expected {
				t.Errorf("Hyperlink(%q, %q) = %q, expected %q", tc.url, tc.text, actual, tc.expected)
			}
		})
	}
}

func TestHyperlinkDisabledColors(t *testing.T) {
	defer DisableColors(false)

	DisableColors(true)
	actual := Hyperlink("https://example.com", "Plain Link")
	if actual != "Plain Link" {
		t.Errorf("Hyperlink with colors disabled = %q, expected %q", actual, "Plain Link")
	}

	actualEmpty := Hyperlink("https://example.com", "")
	if actualEmpty != "" {
		t.Errorf("Hyperlink with colors disabled and empty text = %q, expected %q", actualEmpty, "")
	}

	DisableColors(false)
	actualReenabled := Hyperlink("https://example.com", "Link")
	expected := "\x1b]8;;https://example.com\x1b\\Link\x1b]8;;\x1b\\"
	if actualReenabled != expected {
		t.Errorf("Hyperlink after re-enabling colors = %q, expected %q", actualReenabled, expected)
	}
}
