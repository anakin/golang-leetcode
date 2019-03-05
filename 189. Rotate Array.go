package main

func rotate(nums []int, k int) {
	if k == 0 || len(nums) == 0 {
		return
	}
	k = k % len(nums)
	if k == 0 {
		return
	}
	tmp := nums
	nums = append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
	for i := 0; i < len(tmp); i++ {
		tmp[i] = nums[i]
	}
	return
}
