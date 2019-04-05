package main

import "math"

func increasingTriplet(nums []int) bool {
	return helper(nums, 3)
}

func helper(nums []int, k int) bool {
	if len(nums) < k {
		return false
	}
	dp := make([]int, k+1)
	dp[0] = math.MinInt32
	dp[1] = nums[0]
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		for j := maxLen; j >= 0; j-- {
			if nums[i] > dp[j] {
				if j == maxLen {
					maxLen++
					if maxLen == k {
						return true
					}
				}
				dp[j+1] = nums[i]
				break
			}
		}
	}
	return false
}
