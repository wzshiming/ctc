package ctc

// Color color sum
// The foreground slot and the background slot each hold exactly one color mode (standard, 256-color, or true color),
// so combining two different kinds of foreground (or background) colors with |, e.g. ForegroundColor256(n) | ForegroundBright
// or ForegroundRed | ForegroundColor256(n), produces an undefined color; ForegroundBright/BackgroundBright only combine with the standard colors.
type Color uint64

// Color constant
// Detailed breakdown (64 bits total):
//   - Bits 0-7:   Text styles (Bold, Dim, Italic, Underline, Blink, Negative, Hidden, Strike)
//   - Bits 8-9:   Foreground mode (0=none, 1=Standard, 2=Color256, 3=TrueColor)
//   - Bits 10-33: Foreground color (24 bits: Standard 3bit / Color256 8bit / TrueColor 24bit)
//   - Bits 34-35: Background mode (0=none, 1=Standard, 2=Color256, 3=TrueColor)
//   - Bits 36-59: Background color (24 bits: Standard 3bit / Color256 8bit / TrueColor 24bit)
const (
	Reset Color = 0 // \033[0m

	// Text styles (bits 0-7) - ordered by ANSI escape code sequence
	Bold      Color = 1 << 0 // \033[1m
	Dim       Color = 1 << 1 // \033[2m
	Italic    Color = 1 << 2 // \033[3m
	Underline Color = 1 << 3 // \033[4m
	Blink     Color = 1 << 4 // \033[5m
	Negative  Color = 1 << 5 // \033[7m
	Hidden    Color = 1 << 6 // \033[8m
	Strike    Color = 1 << 7 // \033[9m

	// Standard colors (bits 10-13 for fg, 36-39 for bg)
	// Mode = 1 (Standard)
	bright  = 1 << 3             // 0b1000
	black   = 0                  // 0b0000
	red     = 1 << 0             // 0b0001
	green   = 1 << 1             // 0b0010
	yellow  = red | green        // 0b0011
	blue    = 1 << 2             // 0b0100
	magenta = red | blue         // 0b0101
	cyan    = green | blue       // 0b0110
	white   = red | green | blue // 0b0111

	// Mode constants (2 bits each: 0=none, 1=Standard, 2=Color256, 3=TrueColor)
	modeFgShift = 8    // Foreground mode bits position
	modeBgShift = 34   // Background mode bits position
	modeMask    = 0b11 // 0b11 mask for mode

	// Mode values
	modeStandard  Color = 1 // Standard colors (bits 10-13 for fg, 36-39 for bg)
	modeColor256  Color = 2 // Color256 (8-bit, bits 10-17 or 36-43)
	modeTrueColor Color = 3 // TrueColor (24-bit, bits 10-33 or 36-59)

	// Shift positions for color values
	shiftFg = 10 // Foreground 256 color position
	shiftBg = 36 // Background 256 color position

	// Foreground Standard colors (mode=1, color in bits 10-12)
	applyForegroundStandard       = modeStandard << modeFgShift
	ForegroundBright        Color = applyForegroundStandard | bright<<shiftFg
	ForegroundBlack         Color = applyForegroundStandard | black<<shiftFg
	ForegroundRed           Color = applyForegroundStandard | red<<shiftFg
	ForegroundGreen         Color = applyForegroundStandard | green<<shiftFg
	ForegroundYellow        Color = applyForegroundStandard | yellow<<shiftFg
	ForegroundBlue          Color = applyForegroundStandard | blue<<shiftFg
	ForegroundMagenta       Color = applyForegroundStandard | magenta<<shiftFg
	ForegroundCyan          Color = applyForegroundStandard | cyan<<shiftFg
	ForegroundWhite         Color = applyForegroundStandard | white<<shiftFg

	// Background Standard colors (mode=1, color in bits 36-39)
	applyBackgroundStandard       = modeStandard << modeBgShift
	BackgroundBright        Color = applyBackgroundStandard | bright<<shiftBg
	BackgroundBlack         Color = applyBackgroundStandard | black<<shiftBg
	BackgroundRed           Color = applyBackgroundStandard | red<<shiftBg
	BackgroundGreen         Color = applyBackgroundStandard | green<<shiftBg
	BackgroundYellow        Color = applyBackgroundStandard | yellow<<shiftBg
	BackgroundBlue          Color = applyBackgroundStandard | blue<<shiftBg
	BackgroundMagenta       Color = applyBackgroundStandard | magenta<<shiftBg
	BackgroundCyan          Color = applyBackgroundStandard | cyan<<shiftBg
	BackgroundWhite         Color = applyBackgroundStandard | white<<shiftBg

	// Color256 - 256 color mode (mode=2)
	// 0-7: standard colors
	// 8-15: bright colors
	// 16-231: 6x6x6 RGB cube (216 colors)
	// 232-255: grayscale (24 colors)
	applyForeground256 = modeColor256 << modeFgShift
	applyBackground256 = modeColor256 << modeBgShift
	color256Mask       = 0xFF // 8-bit color value mask

	// TrueColor - 24-bit true color mode (mode=3)
	// RGB values are 0-255 for each channel
	applyForegroundTrueColor = modeTrueColor << modeFgShift
	applyBackgroundTrueColor = modeTrueColor << modeBgShift
)

// ForegroundStandard returns the standard foreground color index (0-15, bit 3 is the bright flag)
// The bool indicates whether there is a standard foreground color
func (c Color) ForegroundStandard() (bool, uint8) {
	if (c>>modeFgShift)&modeMask == modeStandard {
		return true, uint8((c >> shiftFg) & 0xF)
	}
	return false, 0
}

// BackgroundStandard returns the standard background color index (0-15, bit 3 is the bright flag)
// The bool indicates whether there is a standard background color
func (c Color) BackgroundStandard() (bool, uint8) {
	if (c>>modeBgShift)&modeMask == modeStandard {
		return true, uint8((c >> shiftBg) & 0xF)
	}
	return false, 0
}

// ForegroundColor256 returns a 256 color foreground
// color256 is the 256 color code (0-255)
func ForegroundColor256(color256 uint8) Color {
	return applyForeground256 | Color(color256)<<shiftFg
}

// BackgroundColor256 returns a 256 color background
// color256 is the 256 color code (0-255)
func BackgroundColor256(color256 uint8) Color {
	return applyBackground256 | Color(color256)<<shiftBg
}

// ForegroundColor256 returns the 256 color code of the foreground
// The bool indicates whether the foreground is a 256 color
func (c Color) ForegroundColor256() (bool, uint8) {
	if (c>>modeFgShift)&modeMask == modeColor256 {
		return true, uint8((c >> shiftFg) & color256Mask)
	}
	return false, 0
}

// BackgroundColor256 returns the 256 color code of the background
// The bool indicates whether the background is a 256 color
func (c Color) BackgroundColor256() (bool, uint8) {
	if (c>>modeBgShift)&modeMask == modeColor256 {
		return true, uint8((c >> shiftBg) & color256Mask)
	}
	return false, 0
}

// TrueColor - 24-bit RGB color (16.7 million colors)

// ForegroundTrueColor returns a 24-bit true color foreground
// r, g, b are 0-255 for each color channel
func ForegroundTrueColor(r, g, b uint8) Color {
	code := uint64(r)<<16 | uint64(g)<<8 | uint64(b)
	return applyForegroundTrueColor | Color(code)<<shiftFg
}

// BackgroundTrueColor returns a 24-bit true color background
// r, g, b are 0-255 for each color channel
func BackgroundTrueColor(r, g, b uint8) Color {
	code := uint64(r)<<16 | uint64(g)<<8 | uint64(b)
	return applyBackgroundTrueColor | Color(code)<<shiftBg
}

// ForegroundTrueColor returns whether the foreground is a 24-bit true color
// The bool indicates whether the foreground is a true color
func (c Color) ForegroundTrueColor() (bool, uint8, uint8, uint8) {
	if (c>>modeFgShift)&modeMask == modeTrueColor {
		v := uint64(c >> shiftFg)
		r := uint8(v >> 16)
		g := uint8(v >> 8)
		b := uint8(v)
		return true, r, g, b
	}
	return false, 0, 0, 0
}

// BackgroundTrueColor returns whether the background is a 24-bit true color
// The bool indicates whether the background is a true color
func (c Color) BackgroundTrueColor() (bool, uint8, uint8, uint8) {
	if (c>>modeBgShift)&modeMask == modeTrueColor {
		v := uint64(c >> shiftBg)
		r := uint8(v >> 16)
		g := uint8(v >> 8)
		b := uint8(v)
		return true, r, g, b
	}
	return false, 0, 0, 0
}
