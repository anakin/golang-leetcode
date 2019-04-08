package main

/**
二分查找
*/
func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 {
		return 0
	}
	m := len(matrix)
	low, high := matrix[0][0], matrix[m-1][m-1]
	for low < high {
		mid := (low + high) / 2
		cc := helper(matrix, mid)
		if cc < k {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
func helper(matrix [][]int, target int) int {
	c := 0
	i, j := len(matrix)-1, 0
	for i >= 0 && j <= len(matrix)-1 {
		if matrix[i][j] <= target {
			j++
			c += i + 1
		} else {
			i--
		}
	}
	return c
}
