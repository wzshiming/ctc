package ctc

import (
	"bytes"
	"strconv"
	"strings"
	"unsafe"

	_ "github.com/wzshiming/winseq" // Use Unix like Sequences in Windows
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

var pre = []byte("\x1b[0")

// String returns UnixLike markup
func (c Color) String() string {
	b := c.Bytes()
	return *(*string)(unsafe.Pointer(&b))
}

// Bytes returns UnixLike markup
func (c Color) Bytes() []byte {
	// Worst case (fg + bg truecolor 255,255,255 plus all eight styles) is 54 bytes
	s := make([]byte, 0, 64)
	s = append(s, pre...)

	// Foreground colors
	if ok, r, g, b := c.ForegroundTrueColor(); ok {
		// TrueColor foreground: \x1b[38;2;R;G;Bm
		s = appendRGB(s, ";38;2;", r, g, b)
	} else if ok, color := c.ForegroundColor256(); ok {
		// 256 color foreground: \x1b[38;5;Nm
		s = append(s, ";38;5;"...)
		s = strconv.AppendUint(s, uint64(color), 10)
	} else if ok, fgColor := c.ForegroundStandard(); ok {
		// Standard foreground
		if fgColor&0x8 != 0 { // Bright flag in bit 3
			s = appendColor(s, fgColor&0x7, 90)
		} else {
			s = appendColor(s, fgColor, 30)
		}
	}

	// Background colors
	if ok, r, g, b := c.BackgroundTrueColor(); ok {
		// TrueColor background: \x1b[48;2;R;G;Bm
		s = appendRGB(s, ";48;2;", r, g, b)
	} else if ok, color := c.BackgroundColor256(); ok {
		// 256 color background: \x1b[48;5;Nm
		s = append(s, ";48;5;"...)
		s = strconv.AppendUint(s, uint64(color), 10)
	} else if ok, bgColor := c.BackgroundStandard(); ok {
		// Standard background
		if bgColor&0x8 != 0 { // Bright flag in bit 3
			s = appendColor(s, bgColor&0x7, 100)
		} else {
			s = appendColor(s, bgColor, 40)
		}
	}

	// Text styles
	if c&Bold == Bold {
		s = append(s, ';', '1')
	}
	if c&Dim == Dim {
		s = append(s, ';', '2')
	}
	if c&Italic == Italic {
		s = append(s, ';', '3')
	}
	if c&Underline == Underline {
		s = append(s, ';', '4')
	}
	if c&Blink == Blink {
		s = append(s, ';', '5')
	}
	if c&Negative == Negative {
		s = append(s, ';', '7')
	}
	if c&Hidden == Hidden {
		s = append(s, ';', '8')
	}
	if c&Strike == Strike {
		s = append(s, ';', '9')
	}
	s = append(s, 'm')

	return s
}

// Name returns color name
func (c Color) Name() string {
	if c == 0 {
		return "Reset"
	}

	buf := bytes.NewBuffer(nil)
	if ok, fgColor := c.ForegroundStandard(); ok {
		buf.WriteString("Foreground")
		if fgColor&0x8 != 0 {
			buf.WriteString("Bright")
		}
		buf.WriteString(cc[int(fgColor&0x7)])
	} else if ok, v := c.ForegroundColor256(); ok {
		buf.WriteString("ForegroundColor256(")
		buf.WriteString(strconv.Itoa(int(v)))
		buf.WriteString(")")
	} else if ok, r, g, b := c.ForegroundTrueColor(); ok {
		buf.WriteString("ForegroundTrueColor(")
		buf.WriteString(strconv.Itoa(int(r)))
		buf.WriteString(",")
		buf.WriteString(strconv.Itoa(int(g)))
		buf.WriteString(",")
		buf.WriteString(strconv.Itoa(int(b)))
		buf.WriteString(")")
	}

	if ok, bgColor := c.BackgroundStandard(); ok {
		buf.WriteString("Background")
		if bgColor&0x8 != 0 {
			buf.WriteString("Bright")
		}
		buf.WriteString(cc[int(bgColor&0x7)])
	} else if ok, v := c.BackgroundColor256(); ok {
		buf.WriteString("BackgroundColor256(")
		buf.WriteString(strconv.Itoa(int(v)))
		buf.WriteString(")")
	} else if ok, r, g, b := c.BackgroundTrueColor(); ok {
		buf.WriteString("BackgroundTrueColor(")
		buf.WriteString(strconv.Itoa(int(r)))
		buf.WriteString(",")
		buf.WriteString(strconv.Itoa(int(g)))
		buf.WriteString(",")
		buf.WriteString(strconv.Itoa(int(b)))
		buf.WriteString(")")
	}
	if c&Bold == Bold {
		buf.WriteString("Bold")
	}
	if c&Dim == Dim {
		buf.WriteString("Dim")
	}
	if c&Italic == Italic {
		buf.WriteString("Italic")
	}
	if c&Underline == Underline {
		buf.WriteString("Underline")
	}
	if c&Blink == Blink {
		buf.WriteString("Blink")
	}
	if c&Negative == Negative {
		buf.WriteString("Negative")
	}
	if c&Hidden == Hidden {
		buf.WriteString("Hidden")
	}
	if c&Strike == Strike {
		buf.WriteString("Strike")
	}
	return buf.String()
}

// Info returns color details info
func (c Color) Info() string {
	if c == 0 {
		return "Reset"
	}

	ss := []string{}
	if ok, fgColor := c.ForegroundStandard(); ok {
		if fgColor&0x8 != 0 {
			ss = append(ss, "ForegroundBright")
		}
		ss = append(ss, "Foreground"+cc[int(fgColor&0x7)])
	} else if ok, v := c.ForegroundColor256(); ok {
		ss = append(ss, "ForegroundColor256("+strconv.Itoa(int(v))+")")
	} else if ok, r, g, b := c.ForegroundTrueColor(); ok {
		ss = append(ss, "ForegroundTrueColor("+strconv.Itoa(int(r))+","+strconv.Itoa(int(g))+","+strconv.Itoa(int(b))+")")
	}

	if ok, bgColor := c.BackgroundStandard(); ok {
		if bgColor&0x8 != 0 {
			ss = append(ss, "BackgroundBright")
		}
		ss = append(ss, "Background"+cc[int(bgColor&0x7)])
	} else if ok, v := c.BackgroundColor256(); ok {
		ss = append(ss, "BackgroundColor256("+strconv.Itoa(int(v))+")")
	} else if ok, r, g, b := c.BackgroundTrueColor(); ok {
		ss = append(ss, "BackgroundTrueColor("+strconv.Itoa(int(r))+","+strconv.Itoa(int(g))+","+strconv.Itoa(int(b))+")")
	}
	if c&Bold == Bold {
		ss = append(ss, "Bold")
	}
	if c&Dim == Dim {
		ss = append(ss, "Dim")
	}
	if c&Italic == Italic {
		ss = append(ss, "Italic")
	}
	if c&Underline == Underline {
		ss = append(ss, "Underline")
	}
	if c&Blink == Blink {
		ss = append(ss, "Blink")
	}
	if c&Negative == Negative {
		ss = append(ss, "Negative")
	}
	if c&Hidden == Hidden {
		ss = append(ss, "Hidden")
	}
	if c&Strike == Strike {
		ss = append(ss, "Strike")
	}
	return strings.Join(ss, " | ")
}

func appendColor(s []byte, c uint8, off uint8) []byte {
	s = append(s, ';')
	n := uint64(off + c&uint8(white))
	s = strconv.AppendUint(s, n, 10)
	return s
}

func appendRGB(s []byte, prefix string, r, g, b uint8) []byte {
	s = append(s, prefix...)
	s = strconv.AppendUint(s, uint64(r), 10)
	s = append(s, ';')
	s = strconv.AppendUint(s, uint64(g), 10)
	s = append(s, ';')
	s = strconv.AppendUint(s, uint64(b), 10)
	return s
}
