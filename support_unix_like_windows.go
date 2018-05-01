// +build windows

package ctc

import (
	"io"
	"os"
	"sync"
)

var (
	stdout     io.Writer
	stdoutOnce sync.Once
	stderr     io.Writer
	stderrOnce sync.Once
)

func initStdout() {
	stdoutOnce.Do(func() {
		stdout = PipeUnixLike(os.Stdout)
	})
}

func Stdout() io.Writer {
	initStdout()
	return stdout
}

func initStderr() {
	stderrOnce.Do(func() {
		stderr = PipeUnixLike(os.Stderr)
	})
}

func Stderr() io.Writer {
	initStderr()
	return stderr
}
