package ctc

// HTMLMarkup HTML markup
func (c Color) HTMLMarkup() []byte {
	if c&(applyForeground|applyBackground) == 0 {
		return []byte("</span>")
	}
	s := make([]byte, 0, 100)
	s = append(s, []byte("<span style=\"")...)
	if c&applyForeground != 0 {
		s = append(s, []byte("color:#")...)
		s = appendColorHTML(s, uint8(c))
		s = append(s, []byte(";")...)
	}
	if c&applyBackground != 0 {
		s = append(s, []byte("background-color:#")...)
		s = appendColorHTML(s, uint8(c>>4))
		s = append(s, []byte(";")...)
	}
	s = append(s, []byte("\">")...)
	return s
}

func appendColorHTML(s []byte, c uint8) []byte {
	f := []byte("E0")
	z := []byte("00")
	if c&uint8(bright) != 0 {
		f = []byte("FF")
		z = []byte("1F")
	}
	t := uint8(1)
	for i := 0; i != 3; i++ {
		if c&t == 0 {
			s = append(s, f...)
		} else {
			s = append(s, z...)
		}
		t <<= 1
	}
	return s
}
