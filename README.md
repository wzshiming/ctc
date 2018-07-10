# ctc - Console Text Colors

[![Go Report Card](https://goreportcard.com/badge/github.com/wzshiming/ctc)](https://goreportcard.com/report/github.com/wzshiming/ctc)
[![GoDoc](https://godoc.org/github.com/wzshiming/ctc?status.svg)](https://godoc.org/github.com/wzshiming/ctc)
[![GitHub license](https://img.shields.io/github/license/wzshiming/ctc.svg)](https://github.com/wzshiming/ctc/blob/master/LICENSE)
[![cover.run](https://cover.run/go/github.com/wzshiming/ctc.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fwzshiming%2Fctc)

The non-invasive cross-platform terminal color library does not need to modify the Print method

Virtual unix-like environments on Windows

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

## SGR (Select Graphic Rendition)

| Value   | Description       | Behavior                                                          |
| ------: | :---------------- | :---------------------------------------------------------------- |
| 0       | Default           | Returns all attributes to the default state prior to modification |
| 4       | Underline         | Adds underline                                                    |
| 7       | Negative          | Swaps foreground and background colors                            |
| 30~37   | Foreground        | Applies non-bold/bright color to foreground                       |
| 40~47   | Background        | Applies non-bold/bright color to background                       |
| 90~97   | Bright Foreground | Applies bold/bright color to foreground                           |
| 100~107 | Bright Background | Applies bold/bright color to background                           |

## License

Pouch is licensed under the MIT License. See [LICENSE](https://github.com/wzshiming/ctc/blob/master/LICENSE) for the full license text.