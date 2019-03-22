package main

import "sort"

func sortedSquares(A []int) []int {
	for k, v := range A {
		A[k] = v * v
	}
	sort.Ints(A)
	return A
}
