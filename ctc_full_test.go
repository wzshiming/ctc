package ctc

import (
	"strconv"
	"testing"
)

// TestStandardColors tests all Standard colors (foreground and background)
func TestStandardColors(t *testing.T) {
	// Standard foreground colors: 0-7 (black, red, green, yellow, blue, magenta, cyan, white)
	fgColors := []Color{
		ForegroundBlack,
		ForegroundRed,
		ForegroundGreen,
		ForegroundYellow,
		ForegroundBlue,
		ForegroundMagenta,
		ForegroundCyan,
		ForegroundWhite,
	}

	// Standard background colors: 0-7
	bgColors := []Color{
		BackgroundBlack,
		BackgroundRed,
		BackgroundGreen,
		BackgroundYellow,
		BackgroundBlue,
		BackgroundMagenta,
		BackgroundCyan,
		BackgroundWhite,
	}

	// Bright foreground colors: 8-15
	fgBrightColors := []Color{
		ForegroundBright,
		ForegroundBright | ForegroundRed,
		ForegroundBright | ForegroundGreen,
		ForegroundBright | ForegroundYellow,
		ForegroundBright | ForegroundBlue,
		ForegroundBright | ForegroundMagenta,
		ForegroundBright | ForegroundCyan,
		ForegroundBright | ForegroundWhite,
	}

	// Bright background colors: 8-15
	bgBrightColors := []Color{
		BackgroundBright,
		BackgroundBright | BackgroundRed,
		BackgroundBright | BackgroundGreen,
		BackgroundBright | BackgroundYellow,
		BackgroundBright | BackgroundBlue,
		BackgroundBright | BackgroundMagenta,
		BackgroundBright | BackgroundCyan,
		BackgroundBright | BackgroundWhite,
	}

	// Test foreground colors
	for i, c := range fgColors {
		ok, v := c.ForegroundStandard()
		if !ok {
			t.Errorf("ForegroundStandard() should return true for ForegroundBlack, got false")
		}
		if v != uint8(i) {
			t.Errorf("ForegroundStandard() = %d, want %d", v, i)
		}

		// Verify Bytes() contains correct ANSI code
		expectedCode := 30 + i
		if !containsCode(c.Bytes(), expectedCode) {
			t.Errorf("ForegroundColor%d Bytes() should contain %d, got %s", i, expectedCode, c.String())
		}

		// Verify Name() contains expected color name
		expectedNames := []string{"Black", "Red", "Green", "Yellow", "Blue", "Magenta", "Cyan", "White"}
		if !containsSubstring(c.Name(), expectedNames[i]) {
			t.Errorf("ForegroundColor%d Name() should contain %s, got %s", i, expectedNames[i], c.Name())
		}
	}

	// Test background colors
	for i, c := range bgColors {
		ok, v := c.BackgroundStandard()
		if !ok {
			t.Errorf("BackgroundStandard() should return true for BackgroundBlack, got false")
		}
		if v != uint8(i) {
			t.Errorf("BackgroundStandard() = %d, want %d", v, i)
		}

		// Verify Bytes() contains correct ANSI code
		expectedCode := 40 + i
		if !containsCode(c.Bytes(), expectedCode) {
			t.Errorf("BackgroundColor%d Bytes() should contain %d, got %s", i, expectedCode, c.String())
		}

		// Verify Name() contains expected color name
		expectedNames := []string{"Black", "Red", "Green", "Yellow", "Blue", "Magenta", "Cyan", "White"}
		if !containsSubstring(c.Name(), "Background"+expectedNames[i]) {
			t.Errorf("BackgroundColor%d Name() should contain Background%s, got %s", i, expectedNames[i], c.Name())
		}
	}

	// Test bright foreground colors
	for i, c := range fgBrightColors {
		ok, v := c.ForegroundStandard()
		if !ok {
			t.Errorf("ForegroundStandard() should return true for bright color, got false")
		}
		if v != uint8(i|0x8) {
			t.Errorf("ForegroundStandard() = %d, want %d", v, i|0x8)
		}

		// Verify Bytes() contains correct ANSI code (90-97)
		expectedCode := 90 + i
		if !containsCode(c.Bytes(), expectedCode) {
			t.Errorf("ForegroundBright Color%d Bytes() should contain %d, got %s", i, expectedCode, c.String())
		}
	}

	// Test bright background colors
	for i, c := range bgBrightColors {
		ok, v := c.BackgroundStandard()
		if !ok {
			t.Errorf("BackgroundStandard() should return true for bright color, got false")
		}
		if v != uint8(i|0x8) {
			t.Errorf("BackgroundStandard() = %d, want %d", v, i|0x8)
		}

		// Verify Bytes() contains correct ANSI code (100-107)
		expectedCode := 100 + i
		if !containsCode(c.Bytes(), expectedCode) {
			t.Errorf("BackgroundBright Color%d Bytes() should contain %d, got %s", i, expectedCode, c.String())
		}
	}
}

// TestColor256All tests all 256 colors
func TestColor256All(t *testing.T) {
	// Test all 256 color codes
	for color256 := uint8(0); ; color256++ {
		// Test foreground
		fg := ForegroundColor256(color256)
		ok, v := fg.ForegroundColor256()
		if !ok {
			t.Errorf("ForegroundColor256(%d).ForegroundColor256() should return true", color256)
		}
		if v != color256 {
			t.Errorf("ForegroundColor256(%d).ForegroundColor256() = %d, want %d", color256, v, color256)
		}

		// Verify Bytes() contains 38;5;N
		if !containsCode(fg.Bytes(), 38) || !containsCode(fg.Bytes(), 5) {
			t.Errorf("ForegroundColor256(%d).Bytes() should contain 38;5;, got %s", color256, fg.String())
		}
		if !containsSubstring(fg.String(), strconv.Itoa(int(color256))) {
			t.Errorf("ForegroundColor256(%d).String() should contain %d, got %s", color256, color256, fg.String())
		}

		// Verify Name() contains Color256(N)
		if !containsSubstring(fg.Name(), "ForegroundColor256("+strconv.Itoa(int(color256))+")") {
			t.Errorf("ForegroundColor256(%d).Name() should contain Color256(%d), got %s", color256, color256, fg.Name())
		}

		// Verify Info() contains Color256(N)
		if !containsSubstring(fg.Info(), "ForegroundColor256("+strconv.Itoa(int(color256))+")") {
			t.Errorf("ForegroundColor256(%d).Info() should contain Color256(%d), got %s", color256, color256, fg.Info())
		}

		// Test background
		bg := BackgroundColor256(color256)
		ok, v = bg.BackgroundColor256()
		if !ok {
			t.Errorf("BackgroundColor256(%d).BackgroundColor256() should return true", color256)
		}
		if v != color256 {
			t.Errorf("BackgroundColor256(%d).BackgroundColor256() = %d, want %d", color256, v, color256)
		}

		// Verify Bytes() contains 48;5;N
		if !containsCode(bg.Bytes(), 48) || !containsCode(bg.Bytes(), 5) {
			t.Errorf("BackgroundColor256(%d).Bytes() should contain 48;5;, got %s", color256, bg.String())
		}
		if !containsSubstring(bg.String(), strconv.Itoa(int(color256))) {
			t.Errorf("BackgroundColor256(%d).String() should contain %d, got %s", color256, color256, bg.String())
		}

		// Verify Name() contains Color256(N)
		if !containsSubstring(bg.Name(), "BackgroundColor256("+strconv.Itoa(int(color256))+")") {
			t.Errorf("BackgroundColor256(%d).Name() should contain Color256(%d), got %s", color256, color256, bg.Name())
		}

		// Verify Info() contains Color256(N)
		if !containsSubstring(bg.Info(), "BackgroundColor256("+strconv.Itoa(int(color256))+")") {
			t.Errorf("BackgroundColor256(%d).Info() should contain Color256(%d), got %s", color256, color256, bg.Info())
		}

		if color256 == 255 {
			break
		}
	}

	// Test color 255 separately (last color)
	color256 := uint8(255)
	fg := ForegroundColor256(color256)
	ok, v := fg.ForegroundColor256()
	if !ok || v != color256 {
		t.Errorf("ForegroundColor256(255) getter failed")
	}
}

// TestColor256Ranges tests specific Color256 ranges
func TestColor256Ranges(t *testing.T) {
	// Test standard colors (0-7)
	for i := uint8(0); i < 8; i++ {
		c := ForegroundColor256(i)
		ok, v := c.ForegroundColor256()
		if !ok || v != i {
			t.Errorf("ForegroundColor256(%d) getter failed", i)
		}
	}

	// Test bright colors (8-15)
	for i := uint8(8); i < 16; i++ {
		c := ForegroundColor256(i)
		ok, v := c.ForegroundColor256()
		if !ok || v != i {
			t.Errorf("ForegroundColor256(%d) getter failed", i)
		}
	}

	// Test RGB cube (16-231) - sample some values
	// Formula: 16 + 36*r + 6*g + b where r,g,b are 0-5
	testRGBCube := []uint8{16, 17, 21, 22, 51, 52, 196, 197, 231}
	for _, color256 := range testRGBCube {
		c := ForegroundColor256(color256)
		ok, v := c.ForegroundColor256()
		if !ok || v != color256 {
			t.Errorf("ForegroundColor256(%d) getter failed", color256)
		}
	}

	// Test grayscale (232-255)
	for i := uint8(232); ; i++ {
		c := ForegroundColor256(i)
		ok, v := c.ForegroundColor256()
		if !ok || v != i {
			t.Errorf("ForegroundColor256(%d) getter failed", i)
		}
		if i == 255 {
			break
		}
	}
}

// TestColor256WithRGB tests the RGB conversion functions
func TestColor256WithRGB(t *testing.T) {
	// Test RGB to Color256 conversion
	testCases := []struct {
		r, g, b     uint8
		expected256 uint8
	}{
		{0, 0, 0, 16},        // First RGB cube color
		{0, 0, 255, 21},      // Blue max
		{255, 0, 0, 196},     // Red max
		{0, 255, 0, 46},      // Green max
		{255, 255, 255, 231}, // White in RGB cube
	}

	for _, tc := range testCases {
		color256 := RGBToColor256(tc.r, tc.g, tc.b)
		if color256 != tc.expected256 {
			t.Errorf("RGBToColor256(%d, %d, %d) = %d, want %d", tc.r, tc.g, tc.b, color256, tc.expected256)
		}
	}

	// Test Color256 to RGB conversion
	cubeLevels := [6]uint8{0, 95, 135, 175, 215, 255}
	for r := uint8(0); r < 6; r++ {
		for g := uint8(0); g < 6; g++ {
			for b := uint8(0); b < 6; b++ {
				color256 := uint8(16 + 36*r + 6*g + b)
				cr, cg, cb := Color256ToRGB(color256)
				er, eg, eb := cubeLevels[r], cubeLevels[g], cubeLevels[b]
				if cr != er || cg != eg || cb != eb {
					t.Errorf("Color256ToRGB(%d) = (%d, %d, %d), want (%d, %d, %d)", color256, cr, cg, cb, er, eg, eb)
				}
				// Round trip back to the same cube color
				if back := RGBToColor256(cr, cg, cb); back != color256 {
					t.Errorf("RGBToColor256(Color256ToRGB(%d)) = %d, want %d", color256, back, color256)
				}
			}
		}
	}
}

// TestGrayscale tests grayscale colors (232-255)
func TestGrayscale(t *testing.T) {
	// Test grayscale range (232-255)
	for gray := uint8(232); ; gray++ {
		fg := ForegroundColor256(gray)
		ok, v := fg.ForegroundColor256()
		if !ok || v != gray {
			t.Errorf("ForegroundColor256(%d) grayscale getter failed", gray)
		}

		bg := BackgroundColor256(gray)
		ok, v = bg.BackgroundColor256()
		if !ok || v != gray {
			t.Errorf("BackgroundColor256(%d) grayscale getter failed", gray)
		}
		if gray == 255 {
			break
		}
	}
}

// TestColorMatrix tests all combinations of foreground and background
func TestColorMatrix(t *testing.T) {
	// Test a matrix of foreground and background Standard colors
	fgColors := []Color{
		ForegroundBlack,
		ForegroundRed,
		ForegroundGreen,
		ForegroundYellow,
		ForegroundBlue,
		ForegroundMagenta,
		ForegroundCyan,
		ForegroundWhite,
	}

	bgColors := []Color{
		BackgroundBlack,
		BackgroundRed,
		BackgroundGreen,
		BackgroundYellow,
		BackgroundBlue,
		BackgroundMagenta,
		BackgroundCyan,
		BackgroundWhite,
	}

	for _, fg := range fgColors {
		for _, bg := range bgColors {
			combined := fg | bg

			// Verify foreground getter
			ok, fgVal := combined.ForegroundStandard()
			if !ok {
				t.Errorf("ForegroundStandard() should return true for combined color")
			}
			_ = fgVal

			// Verify background getter
			ok, bgVal := combined.BackgroundStandard()
			if !ok {
				t.Errorf("BackgroundStandard() should return true for combined color")
			}
			_ = bgVal

			// Verify Bytes() produces valid output
			b := combined.Bytes()
			if len(b) == 0 {
				t.Errorf("Bytes() should not be empty for combined color")
			}

			// Verify String() produces valid output
			s := combined.String()
			if len(s) == 0 {
				t.Errorf("String() should not be empty for combined color")
			}

			// Verify Name() produces valid output
			n := combined.Name()
			if len(n) == 0 {
				t.Errorf("Name() should not be empty for combined color")
			}

			// Verify Info() produces valid output
			i := combined.Info()
			if len(i) == 0 {
				t.Errorf("Info() should not be empty for combined color")
			}
		}
	}
}

// TestColor256Matrix tests all combinations of foreground and background Color256
func TestColor256Matrix(t *testing.T) {
	// Test a sample matrix of Color256 colors (not all 65536 combinations)
	// Test a few representative values
	sampleColors := []uint8{0, 7, 15, 16, 127, 231, 232, 255}

	for _, fg := range sampleColors {
		for _, bg := range sampleColors {
			fgColor := ForegroundColor256(fg)
			bgColor := BackgroundColor256(bg)
			combined := fgColor | bgColor

			// Verify foreground getter
			ok, fgVal := combined.ForegroundColor256()
			if !ok || fgVal != fg {
				t.Errorf("ForegroundColor256() = %d, want %d", fgVal, fg)
			}

			// Verify background getter
			ok, bgVal := combined.BackgroundColor256()
			if !ok || bgVal != bg {
				t.Errorf("BackgroundColor256() = %d, want %d", bgVal, bg)
			}

			// Verify Bytes() produces valid output
			b := combined.Bytes()
			if len(b) == 0 {
				t.Errorf("Bytes() should not be empty")
			}

			// Verify String() produces valid output
			s := combined.String()
			if len(s) == 0 {
				t.Errorf("String() should not be empty")
			}
		}
	}
}

// TestMixedColors tests mixing Standard and Color256
func TestMixedColors(t *testing.T) {
	// Standard foreground + Color256 background
	fg := ForegroundRed
	bg := BackgroundColor256(200)
	combined := fg | bg

	ok, _ := combined.ForegroundStandard()
	if !ok {
		t.Errorf("ForegroundStandard() should return true")
	}

	ok, _ = combined.BackgroundColor256()
	if !ok {
		t.Errorf("BackgroundColor256() should return true")
	}

	// Color256 foreground + Standard background
	fg = ForegroundColor256(200)
	bg = BackgroundRed
	combined = fg | bg

	ok, _ = combined.ForegroundColor256()
	if !ok {
		t.Errorf("ForegroundColor256() should return true")
	}

	ok, _ = combined.BackgroundStandard()
	if !ok {
		t.Errorf("BackgroundStandard() should return true")
	}
}

// TestReset tests the Reset color
func TestReset(t *testing.T) {
	if Reset.String() != "\x1b[0m" {
		t.Errorf("Reset.String() = %s, want \\x1b[0m", Reset.String())
	}

	if Reset.Name() != "Reset" {
		t.Errorf("Reset.Name() = %s, want Reset", Reset.Name())
	}

	if Reset.Info() != "Reset" {
		t.Errorf("Reset.Info() = %s, want Reset", Reset.Info())
	}
}

// TestTextStyles tests text style constants
func TestTextStyles(t *testing.T) {
	styles := []struct {
		color Color
		name  string
		code  string
	}{
		{Bold, "Bold", "1"},
		{Dim, "Dim", "2"},
		{Italic, "Italic", "3"},
		{Underline, "Underline", "4"},
		{Blink, "Blink", "5"},
		{Negative, "Negative", "7"},
		{Hidden, "Hidden", "8"},
		{Strike, "Strike", "9"},
	}

	for _, s := range styles {
		// Verify Bytes() contains the style code
		if !containsSubstring(string(s.color.Bytes()), s.code) {
			t.Errorf("%s.Bytes() should contain %s, got %s", s.name, s.code, s.color.String())
		}

		// Verify Name() contains the style name
		if !containsSubstring(s.color.Name(), s.name) {
			t.Errorf("%s.Name() should contain %s, got %s", s.name, s.name, s.color.Name())
		}

		// Verify Info() contains the style name
		if !containsSubstring(s.color.Info(), s.name) {
			t.Errorf("%s.Info() should contain %s, got %s", s.name, s.name, s.color.Info())
		}
	}
}

// TestCombinedStyles tests combining multiple text styles
func TestCombinedStyles(t *testing.T) {
	combined := Bold | Italic | Underline

	// Verify all styles are present
	if combined&Bold != Bold {
		t.Errorf("Combined color should have Bold")
	}
	if combined&Italic != Italic {
		t.Errorf("Combined color should have Italic")
	}
	if combined&Underline != Underline {
		t.Errorf("Combined color should have Underline")
	}

	// Verify Bytes() contains all style codes
	b := string(combined.Bytes())
	if !containsSubstring(b, "1") || !containsSubstring(b, "3") || !containsSubstring(b, "4") {
		t.Errorf("Combined.Bytes() should contain 1, 3, 4, got %s", combined.String())
	}
}

// TestColorWithStyles tests colors combined with text styles
func TestColorWithStyles(t *testing.T) {
	color := ForegroundRed | Bold | Underline

	// Verify color is present
	ok, v := color.ForegroundStandard()
	if !ok || v != 1 { // red = 1
		t.Errorf("ForegroundStandard() = %d, want 1", v)
	}

	// Verify styles are present
	if color&Bold != Bold {
		t.Errorf("Color should have Bold")
	}
	if color&Underline != Underline {
		t.Errorf("Color should have Underline")
	}

	// Verify Bytes() contains both color and style codes
	b := string(color.Bytes())
	if !containsSubstring(b, "31") { // red foreground
		t.Errorf("Bytes() should contain 31, got %s", color.String())
	}
	if !containsSubstring(b, "1") { // bold
		t.Errorf("Bytes() should contain 1, got %s", color.String())
	}
	if !containsSubstring(b, "4") { // underline
		t.Errorf("Bytes() should contain 4, got %s", color.String())
	}
}

// TestFieldIsolation verifies foreground and background fields never overlap
func TestFieldIsolation(t *testing.T) {
	// Regression: fg TrueColor high bits must not leak into the bg mode field
	fg := ForegroundTrueColor(255, 255, 255)
	if ok, _, _, _ := fg.BackgroundTrueColor(); ok {
		t.Errorf("ForegroundTrueColor should not report a background: %s", fg.String())
	}
	if ok, _ := fg.BackgroundColor256(); ok {
		t.Errorf("ForegroundTrueColor should not report a 256 color background: %s", fg.String())
	}
	if ok, _ := fg.BackgroundStandard(); ok {
		t.Errorf("ForegroundTrueColor should not report a standard background: %s", fg.String())
	}

	bg := BackgroundTrueColor(255, 255, 255)
	if ok, _, _, _ := bg.ForegroundTrueColor(); ok {
		t.Errorf("BackgroundTrueColor should not report a foreground: %s", bg.String())
	}
	if ok, _ := bg.ForegroundColor256(); ok {
		t.Errorf("BackgroundTrueColor should not report a 256 color foreground: %s", bg.String())
	}
	if ok, _ := bg.ForegroundStandard(); ok {
		t.Errorf("BackgroundTrueColor should not report a standard foreground: %s", bg.String())
	}
}

// Helper function to check if bytes contain a specific ANSI code
func containsCode(b []byte, code int) bool {
	s := string(b)
	return containsSubstring(s, strconv.Itoa(code))
}

// Helper function to check if a string contains a substring
func containsSubstring(s, substr string) bool {
	if len(s) == 0 || len(substr) == 0 {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
