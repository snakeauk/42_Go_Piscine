package piscine

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return (count)
}

func isValid(base string) bool {
	count := 0
	var check string
	for _, c := range base {
		if c == '+' || c == '-' {
			return (false)
		}
		for _, d := range check {
			if c == d {
				return (false)
			}
		}
		check += string(c)
		count++
	}
	if count < 2 {
		return (false)
	}
	return (true)
}

func isBase(base string, c rune) bool {
	for _, b := range base {
		if b == c {
			return (true)
		}
	}
	return (false)
}

func AtoiBase(s string, base string) int {
	if !isValid(base) {
		return (0)
	}
	runes := []rune(s)
	ret := 0
	sign := 1
	len := StrLen(s)
	base_len := StrLen(base)
	index := 0
	if runes[index] == '+' || runes[index] == '-' {
		if runes[index] == '-' {
			sign *= -1
		}
		index++
	}
	for ; index < len; index++ {
		if isBase(base, runes[index]) {
			for j := 0; j < base_len; j++ {
				if s[index] == base[j] {
					ret = ret*base_len + j
				}
			}
		} else {
			return (0)
		}
	}
	return (sign * ret)
}
