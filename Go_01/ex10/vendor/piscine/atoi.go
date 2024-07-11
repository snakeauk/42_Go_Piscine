package piscine

func StrLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return (count)
}

func Atoi(s string) int {
	index := 0
	sign := 1
	ret := 0
	len := StrLen(s)
	if s[index] == '+' || s[index] == '-' {
		if s[index] == '-' {
			sign *= -1
		}
		index++
	}
	for ; index < len; index++ {
		if s[index] >= '0' && s[index] <= '9' {
			ret = ret*10 + int(s[index]-'0')
		} else {
			ret = 0
			break
		}
	}
	return (sign * ret)
}
