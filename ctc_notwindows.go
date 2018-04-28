// +build !windows

package ctc

func init() {
	Style = LikeUnix
}

func (c Color) applyWindows() {
	return
}
