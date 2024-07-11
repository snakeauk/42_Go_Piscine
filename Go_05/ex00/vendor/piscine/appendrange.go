package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var ret []int
	for ; min < max; min++ {
		ret = append(ret, min)
	}
	return (ret)

}
