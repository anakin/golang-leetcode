package main

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
