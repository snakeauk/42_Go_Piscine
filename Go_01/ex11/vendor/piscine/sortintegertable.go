package piscine

func TableSize(table []int) int {
	size := 0
	for range table {
		size++
	}
	return (size)
}

func SortIntegerTable(table []int) {
	size := TableSize(table)
	for i := 0; i < size-1; i++ {
		for j := i; j < size; j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
