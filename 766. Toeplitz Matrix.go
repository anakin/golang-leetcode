package main

func isToeplitzMatrix(matrix [][]int) bool {
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if r > 0 && c > 0 && matrix[r-1][c-1] != matrix[r][c] {
				return false
			}
		}
	}
	return true
}
