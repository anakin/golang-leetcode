package main

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)

	for _, v := range nums1 {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	ret := []int{}

	for _, v := range nums2 {
		if val, ok := m[v]; val > 0 && ok {
			ret = append(ret, v)
			m[v]--
		}
	}
	return ret
}
