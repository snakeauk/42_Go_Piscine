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

func ArrLen(str []string) int {
	count := 0
	for range str {
		count++
	}
	return (count)
}

func RevParams() {
	for i := ArrLen(os.Args) - 1; i > 0; i-- {
		PrintStr(os.Args[i])
		ft.PrintRune('\n')
	}
}
