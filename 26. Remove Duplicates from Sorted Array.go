package main

func removeDuplicates(nums []int) int {
	for i, j := 0, 1; i < j && j < len(nums); {
		if nums[j] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			j--
			i--
		}
		j++
		i++
	}
	return len(nums)
}
