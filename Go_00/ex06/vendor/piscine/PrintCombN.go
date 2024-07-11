package piscine

import "ft"

func PrintString(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
}

func printCombHelper(n int, str string, start rune) {
	if n == 0 {
		PrintString(str)
		if (str[0] - '0') != byte(10-len(str)) {
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
		return
	}
	for i := start; i <= '9'; i++ {
		printCombHelper(n-1, str+string(i), i+1)
	}
}

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	printCombHelper(n, "", '0')
	ft.PrintRune('\n')
}
