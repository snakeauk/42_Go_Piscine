package piscine

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func Index(s string, toFind string) int {
	len := StrLen(s)
	len_toFind := StrLen(toFind)
	if len_toFind == 0 {
		return (-1)
	}
	for i := 0; i <= len-len_toFind; i++ {
		flag := true
		for j := 0; j < len_toFind; j++ {
			if s[i+j] != toFind[j] {
				flag = false
				break
			}
		}
		if flag {
			return (i)
		}
	}
	return (-1)
}
