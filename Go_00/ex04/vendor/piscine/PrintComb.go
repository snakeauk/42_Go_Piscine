package piscine

import "ft"

func PrintComb() {
	for c3 := '0'; c3 <= '9'; c3++ {
		for c2 := '0'; c2 <= '9'; c2++ {
			for c1 := '0'; c1 <= '9'; c1++ {
				if c3 < c2 {
					if c2 < c1 {
						ft.PrintRune(c3)
						ft.PrintRune(c2)
						ft.PrintRune(c1)
						if !(c3 == '7' && c2 == '8' && c1 == '9') {
							ft.PrintRune(',')
							ft.PrintRune(' ')
						}

					}
				}
			}
		}
	}
	ft.PrintRune('\n')
}
