// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/wzshiming/ctc"
)

func plist(beg, end, step ctc.Color) {
	for c := beg; c <= end; c += step {
		a1 := c.Name()
		a2 := c.Info()

		if a1 == a2 {
			fmt.Println("\n//", a1, "existing")
			fmt.Println("//", a1, "=", a2)
		} else {
			fmt.Println("\n//", a1, a2)
			fmt.Println(a1, "=", a2)
		}
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

	out, err := os.OpenFile("./ctc_extra.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	os.Stdout = out

	fmt.Println(`// Code generated; DO NOT EDIT.
//go:generate go run ctc_extra_gen.go
//go:generate go fmt ctc_extra.go

package ctc

const (`)
	for _, v := range data {
		plist(v[0], v[1], v[2])
	}
	fmt.Println(`
)
`)
}
