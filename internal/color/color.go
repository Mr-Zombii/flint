package color

import (
	"fmt"
	"os"
	"runtime"
)

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

var Enabled = isTerminal(os.Stdout.Fd())

func isTerminal(_ uintptr) bool {
	switch runtime.GOOS {
	case "windows":
		return true
	default:
		fi, err := os.Stdout.Stat()
		if err != nil {
			return false
		}
		return (fi.Mode() & os.ModeCharDevice) != 0
	}
}

func Color(text, code string) string {
	if !Enabled {
		return text
	}
	return fmt.Sprintf("%s%s%s", code, text, Reset)
}

func RedText(text string) string {
	return Color(text, Red)
}

func GreenText(text string) string {
	return Color(text, Green)
}

func YellowText(text string) string {
	return Color(text, Yellow)
}

func BlueText(text string) string {
	return Color(text, Blue)
}

func CyanText(text string) string {
	return Color(text, Cyan)
}

func BoldText(text string) string {
	return Color(text, Bold)
}
