package main

/**
trick:先计算总和，然后判断什么时候等于一半
*/
func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum == nums[0] {
		return 0
	}
	half := 0
	for i := 1; i < len(nums); i++ {
		half += nums[i-1]
		if sum-nums[i]-half == half {
			return i
		}
	}
	return -1
}
