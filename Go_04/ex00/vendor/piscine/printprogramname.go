package piscine

import (
	"ft"
	"os"
)

func PrintStr(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
}

func PrintProgramName() {
	PrintStr(os.Args[0])
}
