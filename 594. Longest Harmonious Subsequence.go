package main

/**
用map保存元素出现次数
如果比当前元素大1的元素也出现，则比较和的最大值
*/

func findLHS(nums []int) int {

	ret := 0
	if len(nums) == 0 {
		return ret
	}
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {
		if val, ok := m[k+1]; ok {
			if v+val > ret {
				ret = v + val
			}
		}
	}
	return ret
}
