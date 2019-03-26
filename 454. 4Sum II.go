package main

/**
思路类似两个数相加
*/
func fourSumCount(A []int, B []int, C []int, D []int) int {
	maps := make(map[int]int)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			maps[A[i]+B[j]]++
		}
	}
	ret := 0
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			ret += maps[-1*(C[i]+D[j])]
		}
	}
	return ret
}
