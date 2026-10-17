package main

import (
	"fmt"

	"github.com/wzshiming/ctc"
)

func main() {
	// Show all 256 colors
	// 0-15: Standard and bright colors
	// 16-231: 6x6x6 RGB cube (216 colors)
	// 232-255: Grayscale (24 colors)

	if !ctc.IsTerminal256Color() {
		fmt.Println("Warning: this terminal may not support 256 colors")
		fmt.Println()
	}

	fmt.Println("=== Standard and Bright Colors (0-15) ===")
	for i := uint8(0); i < 16; i++ {
		fmt.Print(ctc.ForegroundColor256(i), i, " ", ctc.Reset)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("=== RGB Cube (16-231) ===")
	fmt.Println("Each row has the same Red value (0-5), columns have same Green (0-5), and values within cell show Blue (0-5)")
	fmt.Println()

	for r := uint8(0); r < 6; r++ {
		for g := uint8(0); g < 6; g++ {
			for b := uint8(0); b < 6; b++ {
				color := ctc.ForegroundColor256(16 + 36*r + 6*g + b)
				fmt.Print(color, 16+36*r+6*g+b, ctc.Reset)
			}
			fmt.Println()
		}
		fmt.Println()
	}

	fmt.Println()

	// Show background colors
	fmt.Println("=== Background RGB Cube ===")
	for r := uint8(0); r < 6; r++ {
		for g := uint8(0); g < 6; g++ {
			for b := uint8(0); b < 6; b++ {
				color := ctc.BackgroundColor256(16+36*r+6*g+b) | ctc.ForegroundBlack
				fmt.Print(color, 16+36*r+6*g+b, ctc.Reset)
			}
			fmt.Println()
		}
		fmt.Println()
	}

	fmt.Println("=== Grayscale (232-255) ===")
	for i := 232; i < 256; i++ {
		fmt.Print(ctc.BackgroundColor256(uint8(i)), " ", ctc.Reset)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("=== RGB to Color256 Conversion ===")
	samples := [][3]uint8{
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
		{255, 128, 0},
		{128, 64, 200},
	}
	for _, s := range samples {
		code := ctc.RGBToColor256(s[0], s[1], s[2])
		r, g, b := ctc.Color256ToRGB(code)
		fmt.Printf("RGB(%3d,%3d,%3d) -> %s color256(%d) = RGB(%3d,%3d,%3d) %s\n",
			s[0], s[1], s[2], ctc.ForegroundColor256(code), code, r, g, b, ctc.Reset)
	}
}
