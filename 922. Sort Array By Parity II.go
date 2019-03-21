package main

func sortArrayByParityII(A []int) []int {
	if len(A) == 0 {
		return []int{}
	}
	ret := make([]int, len(A))
	a, b := 0, 1
	for _, v := range A {
		if v%2 == 0 {
			ret[a] = v
			a += 2
		} else {
			ret[b] = v
			b += 2
		}
	}
	return ret
}
