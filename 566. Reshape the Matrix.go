package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	ret := make([][]int, r)
	for k := range ret {
		ret[k] = make([]int, c)
	}
	if len(nums) == 0 || r*c != len(nums)*len(nums[0]) {
		return nums
	}
	rows, cols := 0, 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			ret[rows][cols] = nums[i][j]
			cols++
			if cols == c {
				rows++
				cols = 0
			}
		}
	}
	return ret
}
