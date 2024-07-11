package piscine

import "ft"

func IsValid(queen_positions []int, row int, col int) bool {
	for i := 0; i < row; i++ {
		if queen_positions[i] == col || queen_positions[i]-col == row-i || col-queen_positions[i] == row-i {
			return (false)
		}
	}
	return (true)
}

func PrintQueenPositions(queen_positions []int) {
	for i := 0; i < 8; i++ {
		ft.PrintRune(rune(queen_positions[i]) + '0')
	}
	ft.PrintRune('\n')
}

func FindQueenPatterns(queen_positions []int, row int, total_patterns *int) int {
	if row == 8 {
		PrintQueenPositions(queen_positions)
		(*total_patterns)++
		return (1)
	}
	for col := 0; col < 8; col++ {
		if IsValid(queen_positions, row, col) {
			queen_positions[row] = col
			FindQueenPatterns(queen_positions, row+1, total_patterns)
		}
	}
	return (0)
}

func EightQueens() int {
	var queen_positions [8]int
	total_patterns := 0
	FindQueenPatterns(queen_positions[:], 0, &total_patterns)
	return (total_patterns)
}
