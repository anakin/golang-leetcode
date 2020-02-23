package main

/*
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example:

Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
 */

func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	helper(&ans, 0, nums)
	return ans
}

func helper(ans *[][]int, depth int, nums []int) {
	if depth == len(nums) {
		*ans = append(*ans, append([]int{}, nums...))
	}

	used := make(map[int]int)
	for i:=depth; i<len(nums); i++ {
		if _, ok := used[nums[i]]; ok {
			continue
		}

		nums[i], nums[depth] = nums[depth], nums[i]
		helper(ans, depth+1, nums)
		nums[i], nums[depth] = nums[depth], nums[i]

		used[nums[i]] = i
	}
}

