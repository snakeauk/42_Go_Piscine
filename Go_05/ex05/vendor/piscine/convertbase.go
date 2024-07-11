package piscine

func StrLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return (count)
}

func StrRune(s string, toFind rune) int {
	len := StrLen(s)
	for index := 0; index < len; index++ {
		if rune(s[index]) == toFind {
			return (index)
		}
	}
	return (-1)
}

func AtoiBase(s, base string) int {
	index := 0
	sign := 1
	ret := 0
	len := StrLen(s)
	base_len := StrLen(base)
	base_index := 0
	if s[index] == '+' || s[index] == '-' {
		if s[index] == '-' {
			sign *= -1
		}
		index++
	}
	for ; index < len; index++ {
		base_index = StrRune(base, rune(s[index]))
		if base_index >= 0 {
			ret = ret*base_len + base_index
		} else {
			ret = 0
			break
		}
	}
	return (sign * ret)
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

func ItoaBase(nbr int, base string) string {
	base_len := StrLen(base)
	var str string
	var c string
	if nbr < 0 {
		str += "-"
		nbr *= -1
	}
	if nbr >= base_len {
		str += ItoaBase(nbr/base_len, base)
	}
	c = string(base[nbr%base_len])
	return (str + c)
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	num := AtoiBase(nbr, baseFrom)
	ret := ItoaBase(num, baseTo)
	return (ret)
}
