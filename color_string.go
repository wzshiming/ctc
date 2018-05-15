package ctc

import (
	"bytes"
	"strings"
)

var cc = []string{
	"Black",
	"Red",
	"Green",
	"Yellow",
	"Blue",
	"Magenta",
	"Cyan",
	"White",
}

func (c Color) String() string {
	return string(c.UnixLikeMarkup())
}

// Name Color name
func (c Color) Name() string {
	if c&(applyForeground|applyBackground) == 0 {
		return "Reset"
	}

	buf := bytes.NewBuffer(nil)
	if c&applyForeground == applyForeground {
		buf.WriteString("Foreground")
		if c&ForegroundBright == ForegroundBright {
			buf.WriteString("Bright")
		}
		buf.WriteString(cc[int(c&foregroundMask&white)])
	}
	if c&applyBackground == applyBackground {
		buf.WriteString("Background")
		if c&BackgroundBright == BackgroundBright {
			buf.WriteString("Bright")
		}
		buf.WriteString(cc[int(((c&backgroundMask)>>4)&white)])
	}
	if c&Underline != 0 {
		buf.WriteString("Underline")
	}
	if c&Negative != 0 {
		buf.WriteString("Negative")
	}
	return buf.String()
}

// Info Color details info
func (c Color) Info() string {
	if c&(applyForeground|applyBackground) == 0 {
		return "Reset"
	}

	ss := []string{}
	if c&applyForeground == applyForeground {
		if c&ForegroundBright == ForegroundBright {
			ss = append(ss, "ForegroundBright")
		}
		ss = append(ss, "Foreground"+cc[int(c&foregroundMask&white)])
	}
	if c&applyBackground == applyBackground {
		if c&BackgroundBright == BackgroundBright {
			ss = append(ss, "BackgroundBright")
		}
		ss = append(ss, "Background"+cc[int(((c&backgroundMask)>>4)&white)])
	}
	if c&Underline != 0 {
		ss = append(ss, "Underline")
	}
	if c&Negative != 0 {
		ss = append(ss, "Negative")
	}
	return strings.Join(ss, " | ")
}
