package netlog

import "testing"

func Test_Red(t *testing.T) {
	Red("red", "one", 1, "two", 2)
}

func Test_Green(t *testing.T) {
	Green("green", "one", 1, "two", 2)
}

func Test_Yellow(t *testing.T) {
	Yellow("yellow", "one", 1, "two", 2)
}

func Test_Blue(t *testing.T) {
	Blue("blue", "one", 1, "two", 2)
}

func Test_Purple(t *testing.T) {
	Purple("purple", "one", 1, "two", 2)
}

func Test_Cyan(t *testing.T) {
	Cyan("cyan", "one", 1, "two", 2)
}

func Test_White(t *testing.T) {
	White("white", "one", 1, "two", 2)
}
