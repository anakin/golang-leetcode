package main

func repeatedNTimes(A []int) int {
	l := len(A) / 2
	m := make(map[int]int)
	for _, v := range A {
		if num, ok := m[v]; ok {
			if num+1 == l {
				return v
			} else {
				m[v]++
			}
		} else {
			m[v] = 1
		}
	}
	return 0
}
