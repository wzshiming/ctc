# (ctc) Console Text Colors

## Support style

- [x] console
  - [x] unix like
  - [x] windows
- [ ] text
  - [ ] html

## SGR (Select Graphic Rendition)

| Value   | Description       | Behavior                                                          |
| ------: | :---------------- | :---------------------------------------------------------------- |
| 0       | Default           | Returns all attributes to the default state prior to modification |
| 4       | Underline         | Adds underline                                                    |
| 7       | Negative          | Swaps foreground and background colors                            |
| 30~37   | Foreground Black  | Applies non-bold/bright color to foreground                       |
| 40~47   | Background Black  | Applies non-bold/bright color to background                       |
| 90~97   | Bright Foreground | Applies bold/bright color to foreground                           |
| 100~107 | Bright Background | Applies bold/bright color to background                           |
