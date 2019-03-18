package main

func findPairs(nums []int, k int) int {
	if len(nums) == 0 || k < 0 {
		return 0
	}
	m := map[int]int{}
	for k, v := range nums {
		m[v] = k
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		if val, ok := m[nums[i]+k]; ok && val != i {
			delete(m, nums[i]+k)
			res++
		}
	}
	return res
}
