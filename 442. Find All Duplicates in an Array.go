package main

func findDuplicates(nums []int) []int {

	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	ret := []int{}
	for k, v := range m {
		if v == 2 {
			ret = append(ret, k)
		}
	}
	return ret
}
