package piscine

import "ft"

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

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return (count)
}
func PrintStr(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
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

func PrintNbrBase(nbr int, base string) {
	var str string
	if !isValid(base) {
		str = "NV"
	} else {
		str = ItoaBase(nbr, base)
	}
	PrintStr(str)
}
