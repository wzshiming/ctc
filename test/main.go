package main

import (
	"fmt"

	"github.com/wzshiming/ctc"
)

func main() {

	fc := func(c ctc.Color) {
		c.Apply()
		fmt.Print(c.String(), " = ", c.Info())
		ctc.Reset.Apply()
		fmt.Printf("\n")
	}
	{
		beg := ctc.ForegroundBlack
		end := ctc.ForegroundWhite | ctc.ForegroundBright + 1
		for i := beg; i != end; i++ {
			fc(i)
		}
	}
	{
		step := ctc.Color(1 << 4)
		beg := ctc.BackgroundBlack
		end := ctc.BackgroundWhite | ctc.BackgroundBright + step
		for i := beg; i != end; i += step {
			fc(i)
		}
	}
	{
		beg := ctc.ForegroundBlack | ctc.BackgroundBlack
		end := ctc.ForegroundWhite | ctc.ForegroundBright | ctc.BackgroundWhite | ctc.BackgroundBright + 1
		for i := beg; i != end; i++ {
			fc(i)
		}
	}
}
