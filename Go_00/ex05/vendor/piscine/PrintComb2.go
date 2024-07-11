package piscine

import "ft"

func PrintInt(nb int) {
	base := 10
	if nb >= base {
		PrintInt(nb / base)
	}
	ft.PrintRune('0' + rune(nb%10))
}

func PrintComb2() {
	nb2_max := 99
	nb1_max := nb2_max - 1
	for nb1 := 0; nb1 <= nb1_max; nb1++ {
		for nb2 := 0; nb2 <= nb2_max; nb2++ {
			if nb1 != nb2 {
				PrintInt(nb1)
				ft.PrintRune(' ')
				PrintInt(nb2)
				if !(nb1 == nb1_max && nb2 == nb2_max) {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
			}
		}
	}
	ft.PrintRune('\n')
}
