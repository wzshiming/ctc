package ctc

const (
	foregroundMask  = bright | white      // 0b0000???? foreground
	backgroundMask  = foregroundMask << 4 // 0b????0000 background
	applyForeground = 1 << 11             // 0b0000100000000000 applyForeground
	applyBackground = 1 << 15             // 0b1000000000000000 ApplyBackground

	bright  = 1 << 3             // 0b1000
	black   = 0                  // 0b?000
	red     = 1 << 0             // 0b?001
	green   = 1 << 1             // 0b?010
	yellow  = red | green        // 0b?011
	blue    = 1 << 2             // 0b?100
	magenta = red | blue         // 0b?101
	cyan    = green | blue       // 0b?110
	white   = red | green | blue // 0b?111
)

const (
	rrMask = red | red<<4
	bbMask = blue | blue<<4
	rbMask = rrMask | bbMask
)

func (c Color) swapRB() Color {
	r := (c & rrMask) << 2
	b := (c & bbMask) >> 2
	return (c & ^Color(rbMask)) | r | b
}

func (c Color) Apply() {
	switch Style {
	case LikeUnix:
		c.applyLikeUnix()
	case Windows:
		c.applyWindows()
	case Markdown:
		c.applyWindows()
	default:
		return
	}
}

func (c Color) Bytes() []byte {
	switch Style {
	case LikeUnix:
		return c.LikeUnixBytes()
	case Windows:
		return nil
	case Markdown:
		return c.MarkdownBytes()
	default:
		return nil
	}
}
