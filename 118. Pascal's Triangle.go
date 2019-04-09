package main

func generate(numRows int) [][]int {
	ret := [][]int{}
	if numRows == 0 {
		return ret
	}
	ret = append(ret, []int{1})
	if numRows == 1 {
		return ret
	}
	for i := 2; i <= numRows; i++ {
		cur := []int{1}
		for j := 1; j < i-1; j++ {
			cur = append(cur, ret[i-2][j-1]+ret[i-2][j])
		}
		cur = append(cur, 1)
		ret = append(ret, cur)
	}
	return ret
}
