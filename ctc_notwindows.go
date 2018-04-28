// +build !windows

package ctc

import (
	"os"
)

func (c Color) apply() {
	os.Stdout.Write(c.Byte())
	return
}
