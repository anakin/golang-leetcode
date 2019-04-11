package main

func transpose(A [][]int) [][]int {
	m := len(A)
	n := len(A[0])
	ans := make([][]int, n)
	for k := range ans {
		ans[k] = make([]int, m)
	}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			ans[c][r] = A[r][c]
		}
	}
	return ans
}
