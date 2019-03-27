package main

/**
先用一个map保存每个元素出现的位置

*/
func findShortestSubArray(nums []int) int {

	m := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = append(m[nums[i]], i)
	}

	degree, ret := 0, len(nums)
	for _, v := range m {
		if degree < len(v) {
			degree = len(v)
			ret = v[len(v)-1] - v[0] + 1
		} else if degree == len(v) {
			if ret > v[len(v)-1]-v[0]+1 {
				ret = v[len(v)-1] - v[0] + 1
			}
		}
	}
	return ret
}
