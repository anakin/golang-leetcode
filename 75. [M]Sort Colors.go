package main

/**
Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:

Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

*/

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
func sortColors(nums []int) {
	l, i, r := 0, 0, len(nums)-1
	for i <= r {
		if nums[i] == 0 {
			swap(nums, i, l)
			i++
			l++
		} else {
			if nums[i] == 1 {
				i++
			} else {
				swap(nums, i, r)
				r--
			}

		}
	}
}
