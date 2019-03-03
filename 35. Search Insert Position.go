package main

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
