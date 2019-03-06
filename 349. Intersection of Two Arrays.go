package main

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, v := range nums1 {
		m[v] = true
	}

	ret := []int{}
	for _, v := range nums2 {
		if val, ok := m[v]; val && ok {
			ret = append(ret, v)
			m[v] = false
		}
	}
	return ret
}
