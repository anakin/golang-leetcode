package main

func largestSumAfterKNegations(A []int, K int) int {
	sum := 0
	for K > 0 {
		small := 0
		for i := 0; i < len(A); i++ {
			if A[i] < A[small] {
				small = i
			}
		}
		A[small] = A[small] * -1
		K--
	}
	for _, v := range A {
		sum += v
	}
	return sum
}
