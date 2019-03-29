package main

/**
定义两个标志位，单调递增和单调递减
扫描一遍，判断是否递增或者递减
*/
func isMonotonic(A []int) bool {
	increasing, decreasing := true, true
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			decreasing = false
		}
		if A[i] < A[i-1] {
			increasing = false
		}
	}
	return increasing || decreasing
}
