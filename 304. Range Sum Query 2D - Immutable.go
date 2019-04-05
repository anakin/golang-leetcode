package main

type NumMatrix struct {
	s [][]int
}

/**
trick是行的方向上，保存累加值
这样计算sum的时候，只有用一次减法就搞定了
*/

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		s := make([][]int, 0)
		return NumMatrix{
			s: s,
		}
	}
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}
	return NumMatrix{
		s: matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row2 >= len(this.s) || col2 >= len(this.s[0]) {
		return 0
	}
	ret := 0
	if col1 == 0 {
		for i := row1; i <= row2; i++ {
			ret += this.s[i][col2]
		}
	} else {
		for i := row1; i <= row2; i++ {
			ret += this.s[i][col2] - this[i][col1-1]
		}
	}
	return ret
}
