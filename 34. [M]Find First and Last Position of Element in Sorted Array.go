package main

/**
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

*/
/**
先用二分查找找到target，再向前后找到边界
*/
func help(nums []int, left, right, t int) int {
	if left >= right {
		if nums[left] != t {
			return -1
		} else {
			return left
		}
	}
	mid := int((left + right) / 2)
	ret := help(nums, left, mid, t)
	if ret != -1 {
		return ret
	}
	return help(nums, mid+1, right, t)
}
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	t := help(nums, 0, len(nums)-1, target)
	if t == -1 {
		return []int{-1, -1}
	}
	left, right := t, t
	for left = t; left >= 0 && nums[left] == target; left-- {
	}
	if left < 0 || nums[left] != target {
		left++
	}
	for right = t; right < len(nums) && nums[right] == target; right++ {

	}
	if right == len(nums) || nums[right] != target {
		right--
	}
	return []int{left, right}
}
