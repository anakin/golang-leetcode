package main

/**
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0
*/

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}
	l := len(nums)
	if target > nums[l-1] {
		return l
	}
	left, right := 0, l-1
	for left <= right {
		m := (left + right) / 2
		if nums[m] > target {
			right = m - 1
			if right >= 0 {
				if nums[right] < target {
					return right + 1
				}
			} else {
				return 0
			}
		} else if nums[m] < target {
			left = m + 1
			if left < l {
				if nums[left] > target {
					return left
				}
			} else {
				return l
			}
		} else {
			return m
		}
	}
	return 0
}
