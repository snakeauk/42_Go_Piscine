package piscine

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return (count)
}

func SplitWhiteSpaces(s string) []string {
	var ret []string
	forward := 0
	for back := 0; back < StrLen(s); back++ {
		if s[back] == ' ' || s[back] == '\t' || s[back] == '\n' {
			if forward != back {
				ret = append(ret, s[forward:back])
				forward = back + 1
			}
		}
	}
	if forward < StrLen(s) {
		ret = append(ret, s[forward:])
	}
	return (ret)
}
