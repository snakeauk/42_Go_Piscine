package piscine

func BasicJoin(elems []string) string {
	var str string
	for _, c := range elems {
		str = str + c
	}
	return (str)
}
