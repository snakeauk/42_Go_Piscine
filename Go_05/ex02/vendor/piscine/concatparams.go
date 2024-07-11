package piscine

func ArrLen(str []string) int {
	count := 0
	for range str {
		count++
	}
	return count
}

func ConcatParams(args []string) string {
	var str []byte
	len := ArrLen(args)
	for i := 0; i < len; i++ {
		str = append(str, args[i]...)
		if i != len-1 {
			str = append(str, '\n')
		}
	}
	return (string(str))
}
