package main

import (
	"math"
)

/**
找到最大的和第二大的，如果最大的大于等于第二大的两倍，就找到了
*/
func dominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	max, sec := math.MinInt32, math.MinInt32
	p := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			sec = max
			max = nums[i]
			p = i
		} else if nums[i] > sec {
			sec = nums[i]
		}
	}
	if max > 2*sec {
		return p
	}
	return -1
}
