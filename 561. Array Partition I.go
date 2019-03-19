package main

import "sort"

//先排序，然后取偶数位置的数相加，就是答案
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ret := 0
	for i := 0; i < len(nums); i += 2 {
		ret += nums[i]
	}
	return ret
}
