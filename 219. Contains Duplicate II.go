package main

//找到的时候判断一下距离
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	m := make(map[int]int)
	for i, v := range nums {
		val, ok := m[nums[i]]
		if ok {
			if val+k >= i {
				return true
			}
		}
		m[v] = i
	}
	return false
}
