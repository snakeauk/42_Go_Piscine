package piscine

func NRune(s string, n int) rune {
	for i, c := range s {
		if i == n-1 {
			return (rune(c))
		}
	}
	return (0)
}
