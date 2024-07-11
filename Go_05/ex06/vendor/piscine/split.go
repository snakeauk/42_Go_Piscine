package piscine

func StrLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return (count)
}

func Split(s, sep string) []string {
	s_len := StrLen(s)
	sep_len := StrLen(sep)
	var start int
	ret := make([]string, 0)
	for i := 0; i <= s_len-sep_len; i++ {
		if s[i:i+sep_len] == sep {
			ret = append(ret, s[start:i])
			start = i + sep_len
		}
	}
	if start <= s_len {
		ret = append(ret, s[start:])
	}
	return ret
}
