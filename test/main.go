package main

import (
	"fmt"

	"github.com/wzshiming/ctc"
)

func plist(beg, end, step ctc.Color) {
	for c := beg; c <= end; c += step {
		fmt.Println(c, c.Name(), " = ", c.Info(), ctc.Reset)
	}
}

func main() {
	data := [...][3]ctc.Color{
		{
			ctc.ForegroundBlack,
			ctc.ForegroundWhite | ctc.ForegroundBright,
			1,
		},
		{
			ctc.BackgroundBlack,
			ctc.BackgroundWhite | ctc.BackgroundBright,
			1 << 4,
		},
		{
			ctc.ForegroundBlack | ctc.BackgroundBlack,
			ctc.ForegroundWhite | ctc.ForegroundBright | ctc.BackgroundWhite | ctc.BackgroundBright,
			1,
		},
	}
	for _, v := range data {
		plist(v[0], v[1], v[2])
	}
}
