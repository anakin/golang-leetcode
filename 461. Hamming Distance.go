package main

func hammingDistance(x int, y int) int {
	res := x ^ y
	ret := 0
	for res > 0 {
		res &= res - 1
		ret++
	}
	return ret
}
