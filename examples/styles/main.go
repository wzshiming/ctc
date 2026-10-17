package main

import (
	"fmt"

	"github.com/wzshiming/ctc"
)

func main() {
	fmt.Println("=== Text Styles ===")
	fmt.Println()

	// Bold
	fmt.Println(ctc.Bold, "Bold", ctc.Reset)

	// Dim
	fmt.Println(ctc.Dim, "Dim", ctc.Reset)

	// Italic
	fmt.Println(ctc.Italic, "Italic", ctc.Reset)

	// Underline
	fmt.Println(ctc.Underline, "Underline", ctc.Reset)

	// Blink
	fmt.Println(ctc.Blink, "Blink", ctc.Reset)

	// Negative
	fmt.Println(ctc.Negative, "Negative", ctc.Reset)

	// Hidden
	fmt.Println(ctc.Hidden, "Hidden", ctc.Reset)

	// Strike
	fmt.Println(ctc.Strike, "Strike", ctc.Reset)

	fmt.Println()
	fmt.Println("=== Combined Styles ===")
	fmt.Println()

	// Multiple styles combined
	boldItalic := ctc.Bold | ctc.Italic
	fmt.Println(boldItalic, "Bold + Italic", ctc.Reset)

	boldUnderline := ctc.Bold | ctc.Underline
	fmt.Println(boldUnderline, "Bold + Underline", ctc.Reset)

	allStyles := ctc.Bold | ctc.Italic | ctc.Underline | ctc.Strike
	fmt.Println(allStyles, "All styles combined", ctc.Reset)

	fmt.Println()
	fmt.Println("=== Styles with Colors ===")
	fmt.Println()

	// Styles with colors
	redBold := ctc.ForegroundRed | ctc.Bold
	fmt.Println(redBold, "Red + Bold", ctc.Reset)

	greenItalic := ctc.ForegroundGreen | ctc.Italic
	fmt.Println(greenItalic, "Green + Italic", ctc.Reset)

	blueUnderline := ctc.ForegroundBlue | ctc.Underline
	fmt.Println(blueUnderline, "Blue + Underline", ctc.Reset)

	yellowStrike := ctc.ForegroundYellow | ctc.Strike
	fmt.Println(yellowStrike, "Yellow + Strike", ctc.Reset)

	whiteBlinkBgRed := ctc.ForegroundWhite | ctc.BackgroundRed | ctc.Blink
	fmt.Println(whiteBlinkBgRed, "White on Red + Blink", ctc.Reset)

	cyanNegative := ctc.ForegroundCyan | ctc.Negative
	fmt.Println(cyanNegative, "Cyan + Negative", ctc.Reset)
}
