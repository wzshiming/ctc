package ctc

// ForEach iterates the standard foreground, background and foreground+background colors in the order ctc_extra.go is generated from
func ForEach(f func(Color)) {
	// Foreground colors (bits 10-13)
	for c := ForegroundBlack; c <= ForegroundWhite|ForegroundBright; c += 1 << 10 {
		f(c)
	}
	// Background colors (bits 36-39)
	for c := BackgroundBlack; c <= BackgroundWhite|BackgroundBright; c += 1 << 36 {
		f(c)
	}
	// All foreground + background combinations
	for bg := BackgroundBlack; bg <= BackgroundWhite|BackgroundBright; bg += 1 << 36 {
		for fg := ForegroundBlack; fg <= ForegroundWhite|ForegroundBright; fg += 1 << 10 {
			f(fg | bg)
		}
	}
}
