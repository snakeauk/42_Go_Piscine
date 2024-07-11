package piscine

func IterativeFactorial(nb int) int {
	ret := 1
	if nb < 0 {
		return (0)
	}
	for ; nb > 1; nb-- {
		ret *= nb
	}
	return (ret)
}
