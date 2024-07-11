package piscine

import (
	"ft"
)

func PrintStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func PrintWordsTables(a []string) {
	for _, s := range a {
		PrintStr(s)
	}
}
