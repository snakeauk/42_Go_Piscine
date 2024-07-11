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
func SortParams() {
	len := ArrLen(os.Args)
	for i := 1; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if os.Args[i] > os.Args[j] {
				os.Args[i], os.Args[j] = os.Args[j], os.Args[i]
			}
		}
	}
	for i := 1; i < len; i++ {
		PrintStr(os.Args[i])
		ft.PrintRune('\n')
	}
}
