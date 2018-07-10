package ctc

import (
	"strconv"

	_ "github.com/wzshiming/winseq" // Use Unix like Sequences in Windows
)

var pre = []byte("\x1b[0")

// UnixLikeMarkup UnixLike markup
func (c Color) UnixLikeMarkup() []byte {
	s := make([]byte, 0, 16)
	s = append(s, pre...)
	if c&applyForeground == applyForeground {
		if c&ForegroundBright == ForegroundBright {
			s = appendColor(s, uint8(c&foregroundMask), 90)
		} else {
			s = appendColor(s, uint8(c&foregroundMask), 30)
		}
	}
	if c&applyBackground == applyBackground {
		if c&BackgroundBright == BackgroundBright {
			s = appendColor(s, uint8(c&backgroundMask>>4), 100)
		} else {
			s = appendColor(s, uint8(c&backgroundMask>>4), 40)
		}
	}
	if c&Underline == Underline {
		s = append(s, ';', '4')
	}
	if c&Negative == Negative {
		s = append(s, ';', '7')
	}
	s = append(s, 'm')

	return s
}

func appendColor(s []byte, c uint8, off uint8) []byte {
	s = append(s, ';')
	n := uint64(off + c&uint8(white))
	s = strconv.AppendUint(s, n, 10)
	return s
}
