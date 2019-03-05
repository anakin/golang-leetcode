package main

//NB
func majorityElement(nums []int) int {
	c := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == c {
			count++
		} else {
			count--
		}
		if count < 1 {
			count = 1
			c = nums[i]
		}
	}
	return c
}
