package main

func smallestRangeI(A []int, K int) int {
	ret := mymax(A) - mymin(A) - 2*K
	if ret > 0 {
		return ret
	}
	return 0
}

func mymin(a []int) int {
	x := a[0]
	for _, v := range a {
		if v < x {
			x = v
		}
	}
	return x
}
func mymax(a []int) int {
	x := a[0]
	for _, v := range a {
		if v > x {
			x = v
		}
	}
	return x
}
