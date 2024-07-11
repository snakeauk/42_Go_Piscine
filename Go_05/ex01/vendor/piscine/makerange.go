package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return (nil)
	}
	ret := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		ret[i] = min + i
	}
	return (ret)
}
