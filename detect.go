package ctc

import (
	"os"
	"strings"
)

// IsTerminal256Color checks if the terminal supports 256 colors
// It only inspects COLORTERM/TERM; it does not check whether stdout is a terminal, NO_COLOR, or Windows signals such as WT_SESSION.
func IsTerminal256Color() bool {
	// TrueColor terminals also support 256 colors
	if IsTerminalTrueColor() {
		return true
	}

	// Check COLORTERM environment variable (most reliable)
	if os.Getenv("COLORTERM") == "256color" {
		return true
	}

	// Check TERM environment variable
	return strings.Contains(os.Getenv("TERM"), "256")
}

// IsTerminalTrueColor checks if the terminal supports 24-bit true color
// It only inspects COLORTERM; it does not check whether stdout is a terminal, NO_COLOR, or Windows signals such as WT_SESSION.
func IsTerminalTrueColor() bool {
	colorterm := os.Getenv("COLORTERM")
	return colorterm == "truecolor" || colorterm == "24bit"
}
