package piscine

func LastRune(s string) rune {
	ret := rune(0)
	for _, c := range s {
		ret = c
	}
	return (ret)
}
