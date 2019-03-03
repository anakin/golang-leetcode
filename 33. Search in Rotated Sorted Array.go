package main

func search(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1

	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[lo] <= nums[mid] {
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

	}
	if nums[lo] == target {
		return lo
	} else {
		return -1
	}
}
