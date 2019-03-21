package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	/**
	ret保存最终结果，l保存临时结果
	*/
	ret, l := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			if l > ret {
				ret = l
			}
			l = 1
		} else {
			l++
		}
		if ret < l {
			ret = l
		}
	}
	return ret
}
