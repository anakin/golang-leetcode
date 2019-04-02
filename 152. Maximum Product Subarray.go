package main

/**
考虑到负负得正的情况，需要同时记录最大值和最小值
*/
func maxProduct(nums []int) int {
	prevMax, prevMin, ret := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == 0 {
			prevMax, prevMin = nums[i], nums[i]
		} else {
			curMax := max(nums[i], max(nums[i]*prevMax, nums[i]*prevMin))
			curMin := min(nums[i], min(nums[i]*prevMax, nums[i]*prevMin))
			prevMax, prevMin = curMax, curMin
		}
		ret = max(prevMax, ret)
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
