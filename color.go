package netlog

import (
	"fmt"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

func apply(color, format string, args ...any) string {
	return color + fmt.Sprintf(format, args...) + reset
}

func Red(format string, args ...any) {
	fmt.Println(apply(red, format, args...))
}

func Green(format string, args ...any) {
	fmt.Println(apply(green, format, args...))
}

func Yellow(format string, args ...any) {
	fmt.Println(apply(yellow, format, args...))
}

func Blue(format string, args ...any) {
	fmt.Println(apply(blue, format, args...))
}

func Purple(format string, args ...any) {
	fmt.Println(apply(purple, format, args...))
}

func Cyan(format string, args ...any) {
	fmt.Println(apply(cyan, format, args...))
}

func White(format string, args ...any) {
	fmt.Println(apply(white, format, args...))
}
