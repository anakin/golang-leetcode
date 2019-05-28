package main

/**
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

func permute(nums []int) [][]int {
	var ret [][]int
	l := len(nums)
	if l == 0 {
		return ret
	}
	helper(nums, 0, l-1, &ret)
	return ret
}

func helper(nums []int, begin, end int, ret *[][]int) {
	if begin == end {
		t := make([]int, len(nums))
		copy(t, nums) //这里一定要copy
		*ret = append(*ret, t)
		return
	}

	for i := begin; i <= end; i++ {
		nums[begin], nums[i] = nums[i], nums[begin]
		helper(nums, begin+1, end, ret)
		nums[begin], nums[i] = nums[i], nums[begin]
	}
}
