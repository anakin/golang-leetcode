package main

/**
滑动窗口
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	count, left, one := 0, 0, 1
	for right := 0; right < len(nums); right++ {
		one *= nums[right]
		for one >= k {
			one /= nums[left]
			left++
		}
		count += right - left + 1
	}
	return count
}
