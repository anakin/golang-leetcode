package main

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	count := 0
	for i < j {
		count += nums[j] - nums[i]
		i++
		j--
	}
	return count
}
