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

	return buf.String()
}

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
	return strings.Join(ss, " | ")
}
