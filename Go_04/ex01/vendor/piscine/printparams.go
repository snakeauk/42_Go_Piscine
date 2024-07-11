package piscine

import (
	"ft"
	"os"
)

func ArrLen(str []string) int {
	count := 0
	for range str {
		count++
	}
	return (count)
}

func PrintStr(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
}

func PrintParams() {
	for i := 1; i < ArrLen(os.Args); i++ {
		PrintStr(os.Args[i])
		ft.PrintRune('\n')
	}
}
