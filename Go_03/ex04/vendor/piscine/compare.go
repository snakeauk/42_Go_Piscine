package piscine

func StrLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}

func Compare(a, b string) int {
	if a < b {
		return (-1)
	}
	if a > b {
		return (1)
	}
	return (0)
}
