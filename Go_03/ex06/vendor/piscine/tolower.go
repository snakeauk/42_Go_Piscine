package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i, c := range runes {
		if c >= 'A' && c <= 'Z' {
			runes[i] = c + ('a' - 'A')
		}
	}
	return (string(runes))
}
