package main

func singleNonDuplicate(nums []int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}
