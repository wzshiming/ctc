# ctc - Console Text Colors

[![GoDoc](https://godoc.org/github.com/wzshiming/ctc?status.svg)](https://godoc.org/github.com/wzshiming/ctc)
[![GitHub license](https://img.shields.io/github/license/wzshiming/ctc.svg)](https://github.com/wzshiming/ctc/blob/master/LICENSE)

The non-invasive cross-platform terminal color library does not need to modify the Print method

Virtual unix-like environments on Windows

Cursor related in [github.com/wzshiming/cursor](https://github.com/wzshiming/cursor)

- [English](https://github.com/wzshiming/ctc/blob/master/README.md)
- [简体中文](https://github.com/wzshiming/ctc/blob/master/README_cn.md)

## Support style

- [x] console
  - [x] unix-like (mac & linux)
  - [x] windows

## example

``` golang
package main

import (
	"fmt"

	"github.com/wzshiming/ctc"
)

func main() {
	// No invasion
	fmt.Println(ctc.BackgroundRed|ctc.ForegroundBlue, "Hello world", ctc.Reset)
}

```

## 256 colors, true color and styles

``` golang
package main

import (
	"fmt"

	"github.com/wzshiming/ctc"
)

func main() {
	// 256 colors
	fmt.Println(ctc.ForegroundColor256(208), "orange", ctc.Reset)

	// True color, falling back to the nearest 256 color
	color := ctc.ForegroundTrueColor(255, 135, 0)
	if !ctc.IsTerminalTrueColor() {
		color = ctc.ForegroundColor256(ctc.RGBToColor256(255, 135, 0))
	}
	fmt.Println(color, "orange", ctc.Reset)

	// Styles
	fmt.Println(ctc.Bold|ctc.Underline|ctc.ForegroundGreen, "styled", ctc.Reset)
}

```

- `ctc.ForegroundColor256(n)`, `ctc.BackgroundColor256(n)`: 256-color palette; 0~7 standard, 8~15 bright, 16~231 6×6×6 cube, 232~255 grayscale
- `ctc.ForegroundTrueColor(r, g, b)`, `ctc.BackgroundTrueColor(r, g, b)`: 24-bit RGB color
- `ctc.Bold`, `ctc.Dim`, `ctc.Italic`, `ctc.Underline`, `ctc.Blink`, `ctc.Negative`, `ctc.Hidden`, `ctc.Strike`: text styles
- `ctc.RGBToColor256(r, g, b)`: nearest color in the 6×6×6 cube (16~231); never returns a grayscale index
- `ctc.Color256ToRGB(n)`: RGB value of a palette index (xterm palette)
- `ctc.IsTerminalTrueColor()`: true if `COLORTERM` is `truecolor` or `24bit`
- `ctc.IsTerminal256Color()`: true if true color is detected, `COLORTERM` is `256color`, or `TERM` contains `256`

Both detection functions only read those environment variables; they do not check for a TTY, `NO_COLOR` or Windows.

Note: the foreground and the background each hold exactly one color mode (standard, 256-color or true color). Combining two different kinds of foreground (or background) colors with `|`, e.g. `ctc.ForegroundColor256(196) | ctc.ForegroundBright` or `ctc.ForegroundRed | ctc.ForegroundColor256(1)`, gives an undefined color. `ForegroundBright` and `BackgroundBright` only combine with the 8 standard colors. Styles, one foreground and one background of any kind combine freely, e.g. `ctc.ForegroundTrueColor(1, 2, 3) | ctc.BackgroundColor256(7) | ctc.Bold` prints `\x1b[0;38;2;1;2;3;48;5;7;1m`.

More examples: [examples/color256](examples/color256), [examples/truecolor](examples/truecolor), [examples/styles](examples/styles).

## SGR (Select Graphic Rendition)

|      Value | Description           | Behavior                                                          |
| ---------: | :-------------------- | :---------------------------------------------------------------- |
|          0 | Default               | Returns all attributes to the default state prior to modification |
|          1 | Bold                  | Increases intensity (bold)                                        |
|          2 | Dim                   | Decreases intensity (faint)                                       |
|          3 | Italic                | Applies italic                                                    |
|          4 | Underline             | Adds underline                                                    |
|          5 | Blink                 | Applies blinking                                                  |
|          7 | Negative              | Swaps foreground and background colors                            |
|          8 | Hidden                | Hides text (conceal)                                              |
|          9 | Strike                | Adds strikethrough                                                |
|      30~37 | Foreground            | Applies non-bold/bright color to foreground                       |
|      40~47 | Background            | Applies non-bold/bright color to background                       |
|      90~97 | Bright Foreground     | Applies bold/bright color to foreground                           |
|    100~107 | Bright Background     | Applies bold/bright color to background                           |
|     38;5;n | Foreground 256-color  | Applies 256-color palette index n (0~255) to foreground           |
|     48;5;n | Background 256-color  | Applies 256-color palette index n (0~255) to background           |
| 38;2;r;g;b | Foreground true color | Applies 24-bit RGB color (r, g, b) to foreground                  |
| 48;2;r;g;b | Background true color | Applies 24-bit RGB color (r, g, b) to background                  |

## License

Licensed under the MIT License. See [LICENSE](https://github.com/wzshiming/ctc/blob/master/LICENSE) for the full license text.
