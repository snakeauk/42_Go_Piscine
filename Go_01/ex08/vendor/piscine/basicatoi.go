package piscine

func BasicAtoi(s string) int {
	ret := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			ret = ret*10 + int(c-'0')
		} else {
			break
		}
	}
	return (ret)
}
