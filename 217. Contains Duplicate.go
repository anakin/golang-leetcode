package main

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	m := make(map[int]bool)
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		}
		m[v] = true
	}
	return false
}
