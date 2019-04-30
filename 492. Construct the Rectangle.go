package main

import "math"

func constructRectangle(area int) []int {
	i := int(math.Sqrt(float64(area)))
	for area%i != 0 {
		i--
	}
	ret := make([]int, 2)
	ret[0] = max(i, area/i)
	ret[1] = min(i, area/i)
	return ret
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
