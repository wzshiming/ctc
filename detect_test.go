package ctc

import (
	"os"
	"testing"
)

// envUnset marks an environment variable that must be absent.
const envUnset = "\x00unset"

func withTermEnv(colorterm, term string, f func()) {
	restore := func(key string) func() {
		old, ok := os.LookupEnv(key)
		return func() {
			if ok {
				os.Setenv(key, old)
			} else {
				os.Unsetenv(key)
			}
		}
	}
	set := func(key, val string) {
		if val == envUnset {
			os.Unsetenv(key)
		} else {
			os.Setenv(key, val)
		}
	}
	defer restore("COLORTERM")()
	defer restore("TERM")()
	set("COLORTERM", colorterm)
	set("TERM", term)
	f()
}

var detectCases = []struct {
	colorterm string
	term      string
	trueColor bool
	color256  bool
}{
	{"truecolor", "", true, true},
	{"24bit", "", true, true},
	{"256color", "", false, true},
	{"", "xterm-256color", false, true},
	{"", "xterm", false, false},
	{envUnset, envUnset, false, false},
	{"truecolor", "dumb", true, true},
	{"yes", "xterm-256color", false, true},
}

func TestIsTerminalTrueColor(t *testing.T) {
	for _, tc := range detectCases {
		withTermEnv(tc.colorterm, tc.term, func() {
			if got := IsTerminalTrueColor(); got != tc.trueColor {
				t.Errorf("IsTerminalTrueColor(COLORTERM=%q, TERM=%q): got %v want %v", tc.colorterm, tc.term, got, tc.trueColor)
			}
		})
	}
}

func TestIsTerminal256Color(t *testing.T) {
	for _, tc := range detectCases {
		withTermEnv(tc.colorterm, tc.term, func() {
			if got := IsTerminal256Color(); got != tc.color256 {
				t.Errorf("IsTerminal256Color(COLORTERM=%q, TERM=%q): got %v want %v", tc.colorterm, tc.term, got, tc.color256)
			}
		})
	}
}
