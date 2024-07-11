package piscine

func ElemCount(s []string) int {
	count := 0
	for range s {
		count++
	}
	return (count)
}
func Join(strs []string, sep string) string {
	elem_i := ElemCount(strs)
	var str string
	for i := 0; i < elem_i; i++ {
		str = str + strs[i]
		if i+1 != elem_i {
			str = str + sep
		}
	}
	return (str)
}
