package main

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	r, c := 0, 0
	m, n := len(matrix), len(matrix[0])
	ret := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		ret[i] = matrix[r][c]
		if (r+c)%2 == 0 {
			if c == n-1 {
				r++
			} else if r == 0 {
				c++
			} else {
				r--
				c++
			}
		} else {
			if r == m-1 {
				c++
			} else if c == 0 {
				r++
			} else {
				r++
				c--
			}
		}
	}
	return ret
}
