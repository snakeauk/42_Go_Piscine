package piscine

func StrLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return (count)
}

func FirstRune(s string) rune {
	for i, c := range s {
		if i == 0 {
			return (rune(c))
		}
	}
	return (0)
}
