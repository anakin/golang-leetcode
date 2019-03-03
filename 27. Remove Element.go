package main

func removeElement(nums []int, val int) int {

	if nums == nil {
		return 0
	}
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
		i++
	}
	return len(nums)

}
