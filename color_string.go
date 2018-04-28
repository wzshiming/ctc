package ctc

import (
	"bytes"
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
