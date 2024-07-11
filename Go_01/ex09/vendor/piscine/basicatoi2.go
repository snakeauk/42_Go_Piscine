package piscine

func BasicAtoi2(s string) int {
	ret := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			ret = ret*10 + int(c-'0')
		} else if c >= 32 && c <= 126 {
			ret = 0
			break
		} else {
			break
		}
	}
	return (ret)
}
