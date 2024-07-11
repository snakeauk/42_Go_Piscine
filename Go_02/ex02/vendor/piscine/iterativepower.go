package piscine

func IterativePower(nb int, power int) int {
	ret := 1
	if power < 0 {
		return (0)
	}
	for ; power > 0; power-- {
		ret *= nb
	}
	return (ret)
}
