package main

/**
先计算数组中所有偶数的和
遍历操作
如果之前是偶数，从和中减去
执行加操作
如果之后是偶数，再加到和里
*/
func sumEvenAfterQueries(A []int, queries [][]int) []int {

	sum := 0
	for _, v := range A {
		if v%2 == 0 {
			sum += v
		}
	}
	ret := []int{}
	for _, v := range queries {
		if A[v[0]]%2 == 0 {
			sum -= A[v[0]]
		}
		A[v[0]] += v[1]
		if A[v[0]]%2 == 0 {
			sum += A[v[0]]
		}
		ret = append(ret, sum)
	}
	return ret
}
