package main

//classic
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
