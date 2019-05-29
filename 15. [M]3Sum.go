package main

import "sort"

/**
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
func threeSum(nums []int) [][]int {
	//先排序
	sort.Ints(nums)

	res := make([][]int, 0)
	f := func(nums []int, begin int, end int, target int) {
		for begin < end { //从两端向中间遍历
			if nums[begin]+nums[end]+target == 0 {
				r := make([]int, 0)
				r = append(r, nums[begin], nums[end], target)
				res = append(res, r)
				//遇到相等的，就快进
				for begin < end && nums[begin] == nums[begin+1] {
					begin++
				}
				for begin < end && nums[end] == nums[end-1] {
					end--
				}
				begin++
				end--
			} else if (target + nums[begin] + nums[end]) < 0 {
				begin++
			} else {
				end--
			}
		}
	}
	l := len(nums)
	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		f(nums, i+1, l-1, nums[i])
	}
	return res
}
