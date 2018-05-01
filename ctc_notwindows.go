// +build !windows

package ctc

func init() {
	Style = UnixLike
}

func (c Color) applyWindows() {
	return
}
