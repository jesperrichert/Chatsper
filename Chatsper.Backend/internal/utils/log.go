package Log

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

func Info(msg string) {
	fmt.Println(pterm.Gray(msg))
}

func Debug(msg string) {
	fmt.Println(pterm.Blue(msg))
}

func Error(msg string, exit bool) {
	fmt.Println(pterm.Red(msg))
	if exit == true {
		os.Exit(1)
	}
}

func Warn(msg string) {
	fmt.Println(pterm.Yellow(msg))
}

func Message(msg string) {
	fmt.Println(pterm.White(msg))
}
