package main

import (
	"fmt"

	"github.com/wzshiming/ctc"
)

func main() {
	ctc.ForEach(func(c ctc.Color) {
		fmt.Println(c, c.Name(), " = ", c.Info(), ctc.Reset)
	})
}
