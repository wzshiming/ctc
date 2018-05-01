package ctc

import (
	"fmt"
)

func Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(Stdout(), a...)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(Stdout(), format, a...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(Stdout(), a...)
}
