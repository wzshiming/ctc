package ctc

type style uint8

const (
	None style = iota
	LikeUnix
	Windows
)

var Style style

type Color uint32

const (
	Reset Color = 0

	ForegroundBright  Color = applyForeground | bright
	ForegroundBlack   Color = applyForeground | black
	ForegroundRed     Color = applyForeground | red
	ForegroundGreen   Color = applyForeground | green
	ForegroundYellow  Color = applyForeground | yellow
	ForegroundBlue    Color = applyForeground | blue
	ForegroundMagenta Color = applyForeground | magenta
	ForegroundCyan    Color = applyForeground | cyan
	ForegroundWhite   Color = applyForeground | white

	BackgroundBright  Color = applyBackground | bright<<4
	BackgroundBlack   Color = applyBackground | black<<4
	BackgroundRed     Color = applyBackground | red<<4
	BackgroundGreen   Color = applyBackground | green<<4
	BackgroundYellow  Color = applyBackground | yellow<<4
	BackgroundBlue    Color = applyBackground | blue<<4
	BackgroundMagenta Color = applyBackground | magenta<<4
	BackgroundCyan    Color = applyBackground | cyan<<4
	BackgroundWhite   Color = applyBackground | white<<4
)
