// +build !windows

package ctc

import (
	"io"
	"os"
)

func init() {
	Style = UnixLike
}

func (c Color) applyWindows() {
	return
}

func Stdout() io.Writer {
	return os.Stdout
}

func Stderr() io.Writer {
	return os.Stderr
}
