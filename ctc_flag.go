package ctc

const (
	foregroundMask  Color = bright | white      // 0b0000???? foreground
	backgroundMask  Color = foregroundMask << 4 // 0b????0000 background
	applyForeground Color = 1 << 8              // 0b0000000100000000 applyForeground
	applyBackground Color = 1 << 9              // 0b0000001000000000 ApplyBackground

	bright  Color = 1 << 3             // 0b1000
	black   Color = 0                  // 0b?000
	red     Color = 1 << 0             // 0b?001
	green   Color = 1 << 1             // 0b?010
	yellow  Color = red | green        // 0b?011
	blue    Color = 1 << 2             // 0b?100
	magenta Color = red | blue         // 0b?101
	cyan    Color = green | blue       // 0b?110
	white   Color = red | green | blue // 0b?111
)

const (
	rrMask = red | red<<4
	bbMask = blue | blue<<4
	rbMask = rrMask | bbMask
)

func (c Color) swapRB() Color {
	r := (c & rrMask) << 2
	b := (c & bbMask) >> 2
	return (c & ^rbMask) | r | b
}

func (c Color) Apply() {
	c.apply()
}
