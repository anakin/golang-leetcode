package main

/**
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
*/
func maxSubArray(nums []int) int {
	cur := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum >= 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if cur < sum {
			cur = sum
		}
	}
	return cur
}
