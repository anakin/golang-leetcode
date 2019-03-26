package main

/**
简单的二分查找
*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid - 1
			continue
		}
		if nums[mid] < target {
			start = mid + 1
		}
	}
	return -1
}
