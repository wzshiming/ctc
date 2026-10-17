# ctc - 终端文本颜色

[![GoDoc](https://godoc.org/github.com/wzshiming/ctc?status.svg)](https://godoc.org/github.com/wzshiming/ctc)
[![GitHub license](https://img.shields.io/github/license/wzshiming/ctc.svg)](https://github.com/wzshiming/ctc/blob/master/LICENSE)

无侵入的跨平台的终端颜色库不需要修改Print方法

在windows 下靠虚拟成类unix 环境

与光标相关的在 [github.com/wzshiming/cursor](https://github.com/wzshiming/cursor)

- [English](https://github.com/wzshiming/ctc/blob/master/README.md)
- [简体中文](https://github.com/wzshiming/ctc/blob/master/README_cn.md)

## 支持

- [x] 终端
  - [x] 类 unix (mac 和 linux)
  - [x] windows

## 示例

``` golang
package main

import (
	"fmt"

	"github.com/wzshiming/ctc"
)

func main() {
	// 无侵入
	fmt.Println(ctc.BackgroundRed|ctc.ForegroundBlue, "Hello world", ctc.Reset)
}

```

## 256 色、真彩色与样式

``` golang
package main

import (
	"fmt"

	"github.com/wzshiming/ctc"
)

func main() {
	// 256 色
	fmt.Println(ctc.ForegroundColor256(208), "orange", ctc.Reset)

	// 真彩色, 不支持时退回到最接近的 256 色
	color := ctc.ForegroundTrueColor(255, 135, 0)
	if !ctc.IsTerminalTrueColor() {
		color = ctc.ForegroundColor256(ctc.RGBToColor256(255, 135, 0))
	}
	fmt.Println(color, "orange", ctc.Reset)

	// 样式
	fmt.Println(ctc.Bold|ctc.Underline|ctc.ForegroundGreen, "styled", ctc.Reset)
}

```

- `ctc.ForegroundColor256(n)`, `ctc.BackgroundColor256(n)`: 256 色; 0~7 标准色, 8~15 高亮色, 16~231 6×6×6 色立方, 232~255 灰阶
- `ctc.ForegroundTrueColor(r, g, b)`, `ctc.BackgroundTrueColor(r, g, b)`: 24 位 RGB 真彩色
- `ctc.Bold`, `ctc.Dim`, `ctc.Italic`, `ctc.Underline`, `ctc.Blink`, `ctc.Negative`, `ctc.Hidden`, `ctc.Strike`: 文本样式
- `ctc.RGBToColor256(r, g, b)`: 6×6×6 色立方 (16~231) 中最接近的颜色, 不会返回灰阶
- `ctc.Color256ToRGB(n)`: 256 色对应的 RGB 值 (xterm 调色板)
- `ctc.IsTerminalTrueColor()`: `COLORTERM` 为 `truecolor` 或 `24bit` 时返回 true
- `ctc.IsTerminal256Color()`: 支持真彩色, 或 `COLORTERM` 为 `256color`, 或 `TERM` 包含 `256` 时返回 true

两个检测函数只读取上述环境变量, 不检查 TTY、`NO_COLOR` 或 Windows。

注意: 前景和背景各自只能使用一种颜色模式 (标准色、256 色或真彩色)。用 `|` 组合两种不同类型的前景 (或背景) 颜色, 例如 `ctc.ForegroundColor256(196) | ctc.ForegroundBright` 或 `ctc.ForegroundRed | ctc.ForegroundColor256(1)`, 结果颜色是未定义的。`ForegroundBright` 和 `BackgroundBright` 只能与 8 种标准色组合。样式与任意类型的一个前景加一个背景可以自由组合, 例如 `ctc.ForegroundTrueColor(1, 2, 3) | ctc.BackgroundColor256(7) | ctc.Bold` 输出 `\x1b[0;38;2;1;2;3;48;5;7;1m`。

更多示例: [examples/color256](examples/color256)、[examples/truecolor](examples/truecolor)、[examples/styles](examples/styles)。

## SGR (Select Graphic Rendition)

|      Value | Description | Behavior                 |
| ---------: | :---------- | :----------------------- |
|          0 | 恢复默认    | 把所有属性还原到修改之前 |
|          1 | 加粗        | 加粗文字                 |
|          2 | 变暗        | 降低亮度                 |
|          3 | 斜体        | 斜体文字                 |
|          4 | 下划线      | 添加下划线               |
|          5 | 闪烁        | 文字闪烁                 |
|          7 | 交换颜色    | 交换前景和背景的颜色     |
|          8 | 隐藏        | 隐藏文字                 |
|          9 | 删除线      | 添加删除线               |
|      30~37 | 前景        | 非高量的前景颜色         |
|      40~47 | 背景        | 非高量的背景颜色         |
|      90~97 | 高亮前景    | 高量的前景颜色           |
|    100~107 | 高亮背景    | 高量的背景颜色           |
|     38;5;n | 256 色前景  | 前景使用第 n 个 256 色   |
|     48;5;n | 256 色背景  | 背景使用第 n 个 256 色   |
| 38;2;r;g;b | 真彩前景    | 前景使用 24 位 RGB 颜色  |
| 48;2;r;g;b | 真彩背景    | 背景使用 24 位 RGB 颜色  |

## 许可证

软包根据MIT License。有关完整的许可证文本，请参阅[LICENSE](https://github.com/wzshiming/ctc/blob/master/LICENSE).
