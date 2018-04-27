package ctc

type Color uint32

const (
	Reset Color = 0

	BrightForeground  = applyForeground | bright
	BlackForeground   = applyForeground | black
	RedForeground     = applyForeground | red
	GreenForeground   = applyForeground | green
	YellowForeground  = applyForeground | yellow
	BlueForeground    = applyForeground | blue
	MagentaForeground = applyForeground | magenta
	CyanForeground    = applyForeground | cyan
	WhiteForeground   = applyForeground | white

	BrightBackground  = applyBackground | bright<<4
	BlackBackground   = applyBackground | black<<4
	RedBackground     = applyBackground | red<<4
	GreenBackground   = applyBackground | green<<4
	YellowBackground  = applyBackground | yellow<<4
	BlueBackground    = applyBackground | blue<<4
	MagentaBackground = applyBackground | magenta<<4
	CyanBackground    = applyBackground | cyan<<4
	WhiteBackground   = applyBackground | white<<4
)
