package main

/**
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

Example:

Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

func combine(n int, k int) [][]int {
	var res [][]int
	helper(&res, []int{}, 1, n, k)
	return res
}

func helper(res *[][]int, coms []int, start int, n int, k int) {
	if k == 0 {
		tmp := []int{}
		tmp = append(tmp, coms...)
		*res = append(*res, tmp)
		return
	}
	for i := start; i <= n; i++ {
		coms = append(coms, i)
		helper(res, coms, i+1, n, k-1)
		coms = coms[:len(coms)-1]
	}
}
