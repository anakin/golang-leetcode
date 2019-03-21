package main

/**
采用快排分治的思想
*/
func sortArrayByParity(A []int) []int {
	l, r := 0, len(A)-1
	tmp := A[l]
	for l < r {
		for l < r && A[r]%2 == 1 {
			r--
		}
		A[l] = A[r]
		for l < r && A[l]%2 == 0 {
			l++
		}
		A[r] = A[l]
	}
	A[l] = tmp
	return A
}
