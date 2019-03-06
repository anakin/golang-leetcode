package main

func findDuplicate(nums []int) int {
	left, right := 1, len(nums)-1
	mid := left + (right-left)/2
	for left < right {
		c := 0
		mid = left + (right-left)/2
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				c++
			}
		}
		if c > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
