// +build windows

package ctc

import (
	"syscall"
	"unsafe"
)

func init() {
	applyUnixlikeMode()
}

var (
	kernel32                        = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleMode              = kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode              = kernel32.NewProc("SetConsoleMode")
	hStdout                         = uintptr(syscall.Stdout)
	enableVirtualTerminalProcessing = 0x0004
)

func applyUnixlikeMode() error {
	var mode uintptr
	ok, _, err := procGetConsoleMode.Call(hStdout, uintptr(unsafe.Pointer(&mode)))
	if ok != 0 {
		err = nil
	}
	if err != nil {
		return err
	}
	ok, _, err = procSetConsoleMode.Call(hStdout, mode|enableVirtualTerminalProcessing)
	if ok != 0 {
		err = nil
	}
	return err
}
