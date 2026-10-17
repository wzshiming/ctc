package ctc

// standardPalette is the xterm RGB values of the 16 standard colors (0-15)
var standardPalette = [16][3]uint8{
	{0, 0, 0},
	{128, 0, 0},
	{0, 128, 0},
	{128, 128, 0},
	{0, 0, 128},
	{128, 0, 128},
	{0, 128, 128},
	{192, 192, 192},
	{128, 128, 128},
	{255, 0, 0},
	{0, 255, 0},
	{255, 255, 0},
	{0, 0, 255},
	{255, 0, 255},
	{0, 255, 255},
	{255, 255, 255},
}

// Color256ToRGB converts a 256 color code to its xterm RGB values (0-255)
// For standard colors (0-15): returns the standard palette values
// For the RGB cube (16-231): each channel level n (0-5) maps to 0 or n*40+55
// For grayscale (232-255): 232 + n maps to gray value n*10 + 8
func Color256ToRGB(color256 uint8) (r, g, b uint8) {
	if color256 < 16 {
		c := standardPalette[color256]
		return c[0], c[1], c[2]
	}
	if color256 >= 232 {
		gray := (color256-232)*10 + 8
		return gray, gray, gray
	}
	// RGB cube: 16 + 36*r + 6*g + b
	v := color256 - 16
	return cubeLevelToRGB(v / 36), cubeLevelToRGB((v / 6) % 6), cubeLevelToRGB(v % 6)
}

// cubeLevelToRGB converts a cube channel level (0-5) to its xterm RGB value
func cubeLevelToRGB(level uint8) uint8 {
	if level == 0 {
		return 0
	}
	return level*40 + 55
}

// RGBToColor256 converts RGB values (0-255) to a 256 color code
// Each channel is quantized to the nearest cube level, then the color code
// is calculated in the RGB cube range (16-231)
func RGBToColor256(r, g, b uint8) uint8 {
	return 16 + 36*rgbToCubeLevel(r) + 6*rgbToCubeLevel(g) + rgbToCubeLevel(b)
}

// rgbToCubeLevel quantizes an RGB channel value (0-255) to the nearest cube level (0-5)
func rgbToCubeLevel(v uint8) uint8 {
	if v < 48 {
		return 0
	}
	if v < 115 {
		return 1
	}
	return (v - 35) / 40
}
