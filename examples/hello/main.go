package main

import (
	"fmt"

	"github.com/wzshiming/ctc"
)

func main() {
	// No invasion
	fmt.Println(ctc.BackgroundRed|ctc.ForegroundBlue, "Hello world", ctc.Reset)
}
