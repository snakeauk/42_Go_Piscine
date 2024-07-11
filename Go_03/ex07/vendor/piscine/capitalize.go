package piscine

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func Capitalize(s string) string {
	runes := []rune(s)
	maxIndex := StrLen(s)
	for i := 0; i < maxIndex; i++ {
		if i == 0 || runes[i-1] == ' ' {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = runes[i] - ('a' - 'A')
			}
		} else {
			if runes[i] >= 'A' && runes[i] <= 'Z' {
				runes[i] = runes[i] + ('a' - 'A')
			}
		}
	}

	return string(runes)
}
