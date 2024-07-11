package piscine

func Sqrt(nb int) int {
	for n := 0; n <= nb/2; n++ {
		if n*n == nb {
			return (n)
		}
		if n*n > nb {
			break
		}
	}
	return (0)
}
