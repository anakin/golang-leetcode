package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	f := func(nums []int, begin int, end int, target int) {
		for begin < end {
			if nums[begin]+nums[end]+target == 0 {
				r := make([]int, 0)
				r = append(r, nums[begin], nums[end], target)
				res = append(res, r)
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
