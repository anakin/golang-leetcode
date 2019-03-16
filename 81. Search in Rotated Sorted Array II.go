package main

func search(nums []int, target int) bool {
	l := len(nums)
	if l == 0 {
		return false
	}
	left, right := 0, l-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if nums[mid] > nums[right] {
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			right--
		}

	}
	return false
}
