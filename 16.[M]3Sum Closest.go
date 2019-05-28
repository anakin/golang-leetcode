package main

import (
	"math"
	"sort"
)

/**
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

*/

func threeSumClosest(nums []int, target int) int {
	dist := math.MaxInt32
	ret := 0
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	l := len(nums)
	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		begin, end := i+1, l-1
		for begin < end {
			sum := nums[begin] + nums[end] + nums[i]
			if sum < target {
				if target-sum < dist {
					dist = target - sum
					ret = sum
				}
				begin++
			} else if sum > target {
				if sum-target < dist {
					dist = sum - target
					ret = sum
				}
				end--
			} else {
				return sum
			}
		}
	}
	return ret
}
