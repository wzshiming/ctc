// +build windows

package ctc

import (
	"syscall"
	"unsafe"
)

func init() {
	if initScreenInfo != nil {
		Style = Windows
	}
}

func (c Color) applyWindows() {
	if initScreenInfo == nil {
		return
	}
	w := initScreenInfo.WAttributes
	if c&(applyForeground|applyBackground) != 0 {
		w = uint16(c.swapRB() & (backgroundMask | foregroundMask))
	}
	setConsoleTextAttribute(hStdout, w)
	return
}

var (
	kernel32                       = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTextAttribute    = kernel32.NewProc("SetConsoleTextAttribute")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	hStdout                        = uintptr(syscall.Stdout)
	initScreenInfo                 = getConsoleScreenBufferInfo(hStdout)
)

func getConsoleScreenBufferInfo(hStdout uintptr) *consoleScreenBufferInfo {
	var csbi consoleScreenBufferInfo
	if ret, _, _ := procGetConsoleScreenBufferInfo.Call(hStdout, uintptr(unsafe.Pointer(&csbi))); ret == 0 {
		return nil
	}
	return &csbi
}

func setConsoleTextAttribute(hStdout uintptr, wAttributes uint16) bool {
	ret, _, _ := procSetConsoleTextAttribute.Call(
		hStdout,
		uintptr(wAttributes))
	return ret != 0
}

type coord struct {
	X, Y int16
}

type smallRect struct {
	Left, Top, Right, Bottom int16
}

type consoleScreenBufferInfo struct {
	DwSize              coord
	DwCursorPosition    coord
	WAttributes         uint16
	SrWindow            smallRect
	DwMaximumWindowSize coord
}
