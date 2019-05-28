package main

/**
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

/**
用hash保存另一个数的值，遍历的同时判断是否存在
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		c, ok := m[nums[i]]
		if ok {
			r := []int{c, i}
			return r
		} else {
			m[target-nums[i]] = i
		}
	}
	r := []int{0, 0}
	return r
}
