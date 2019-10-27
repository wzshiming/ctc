// +build ignore

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/wzshiming/ctc"
)

func main() {

	buf := bytes.NewBuffer(nil)

	fmt.Fprintln(buf, `// Code generated; DO NOT EDIT.

//go:generate go run ctc_extra_gen.go
//go:generate go fmt ctc_extra.go

package ctc

const (`)
	ctc.ForEach(func(c ctc.Color) {
		a1 := c.Name()
		a2 := c.Info()

		if a1 == a2 {
			fmt.Fprintln(buf, "\n//", a1, "existing")
			fmt.Fprintln(buf, "//", a1, "=", a2)
		} else {
			fmt.Fprintln(buf, "\n//", a1, a2)
			fmt.Fprintln(buf, a1, "=", a2)
		}
	})
	fmt.Fprintln(buf, `
)
`)

	ioutil.WriteFile("./ctc_extra.go", buf.Bytes(), 0600)
}
