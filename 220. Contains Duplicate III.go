package main

/**
双重循环，维持一个滑动窗口
*/

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) == 0 {
		return false
	}
	for i := 1; i < len(nums); i++ {
		j := 0
		if i-k > 0 {
			j = i - k
		}
		for ; j >= 0 && j < i; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}
func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
