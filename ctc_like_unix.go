package ctc

import (
	"os"
	"strconv"
)

func (c Color) LikeUnixBytes() []byte {
	s := make([]byte, 0, 16)
	s = append(s, []byte("\x1b[0")...)
	if c&applyForeground != 0 {
		s = appendColor(s, uint8(c&foregroundMask), 30)
	}
	if c&applyBackground != 0 {
		s = appendColor(s, uint8(c&backgroundMask>>4), 40)
	}
	s = append(s, 'm')

	return s
}

func appendColor(s []byte, c uint8, off uint8) []byte {
	s = append(s, ';')
	n := uint64(off + c&uint8(white))
	s = strconv.AppendUint(s, n, 10)
	if c&uint8(bright) != 0 {
		s = append(s, ';', '1')
	}
	return s
}

func (c Color) applyLikeUnix() {
	os.Stdout.Write(c.LikeUnixBytes())
	return
}
