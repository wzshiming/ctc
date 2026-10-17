package main

import (
	"fmt"

	"github.com/wzshiming/ctc"
)

func main() {
	fmt.Println("TrueColor (24-bit) Examples:")
	fmt.Println("===========================")

	if !ctc.IsTerminalTrueColor() {
		fmt.Println("Warning: this terminal may not support true color, consider falling back to 256 colors with ctc.RGBToColor256")
	}

	// Gradient example
	fmt.Println()
	fmt.Println("Gradient (rainbow):")
	text := "The quick brown fox jumps over the lazy dog! Rainbow gradient makes the world more beautiful!"
	for i := 0; i < len(text); i++ {
		h := float64(i) / float64(len(text)-1) * 360
		r, g, b := hsvToRgb(h, 1, 1)
		fmt.Printf("%s%c", foreground(r, g, b), text[i])
	}
	fmt.Println(ctc.Reset)

	// Gradient with background
	fmt.Println()
	fmt.Println("Color palette:")
	rows, cols := 80, 80
	for y := 0; y < rows; y++ {
		s := 1.0 - float64(y)/float64(rows-1)*0.8
		for x := 0; x < cols; x++ {
			h := float64(x) / float64(cols-1) * 360
			v := 0.3 + float64(y)/float64(rows-1)*0.7
			r, g, b := hsvToRgb(h, s, v)
			ir, ig, ib := 255-r, 255-g, 255-b
			fmt.Printf("%s ", ctc.ForegroundTrueColor(ir, ig, ib)|ctc.BackgroundTrueColor(r, g, b))
		}
		fmt.Println(ctc.Reset)
	}

	// Grayscale
	fmt.Println()
	fmt.Println("Grayscale (TrueColor):")
	for i := 0; i < 80; i++ {
		v := uint8(i * 255 / 79)
		fmt.Printf("%s ", ctc.ForegroundTrueColor(255-v, 255-v, 255-v)|ctc.BackgroundTrueColor(v, v, v))
	}
	fmt.Println(ctc.Reset)

}

// foreground returns a truecolor foreground, falling back to 256 colors
// when the terminal does not support truecolor
func foreground(r, g, b uint8) ctc.Color {
	if ctc.IsTerminalTrueColor() {
		return ctc.ForegroundTrueColor(r, g, b)
	}
	return ctc.ForegroundColor256(ctc.RGBToColor256(r, g, b))
}

// hsvToRgb converts HSV to RGB (h: 0-360, s: 0-1, v: 0-1)
func hsvToRgb(h, s, v float64) (r, g, b uint8) {
	i := int(h / 60)
	f := h/60 - float64(i)
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)

	switch i % 6 {
	case 0:
		return uint8(v * 255), uint8(t * 255), uint8(p * 255)
	case 1:
		return uint8(q * 255), uint8(v * 255), uint8(p * 255)
	case 2:
		return uint8(p * 255), uint8(v * 255), uint8(t * 255)
	case 3:
		return uint8(p * 255), uint8(q * 255), uint8(v * 255)
	case 4:
		return uint8(t * 255), uint8(p * 255), uint8(v * 255)
	default:
		return uint8(v * 255), uint8(p * 255), uint8(q * 255)
	}
}
