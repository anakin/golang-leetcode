package main

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
