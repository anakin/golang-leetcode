package main

import (
	"math"
)

func find132pattern(nums []int) bool {
	m := math.MaxInt32
	for j := 0; j < len(nums); j++ {
		m = min(nums[j], m)
		if m == nums[j] {
			continue
		}
		for k := len(nums) - 1; k > j; k-- {
			if m < nums[k] && nums[k] < nums[j] {
				return true
			}
		}
	}
	return false
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

//fastest AC code
//use stack
func find132pattern1(nums []int) bool {
	ell := len(nums)
	if ell < 3 {
		return false
	}

	stack := make([]int, 1, ell)
	stack[0] = nums[ell-1]

	for i, s3 := ell-2, math.MinInt32; i >= 0; i-- {
		if nums[i] < s3 {
			return true
		}

		for j := len(stack) - 1; j >= 0 && stack[j] < nums[i]; j-- {
			s3, stack = stack[j], stack[:j]
		}
		stack = append(stack, nums[i])
	}

	return false
}
