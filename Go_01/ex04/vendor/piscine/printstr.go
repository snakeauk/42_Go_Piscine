package piscine

import "ft"

func PrintStr(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
}
