package main

import (
	"math"
	"sort"
)

/**
用二分查找。对heaters进行sort，然后得到某个house的相对位置。再计算左右的距离。

*/
func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	ret := -1
	for _, house := range houses {
		idx := sort.SearchInts(heaters, house)
		l, r := math.MaxInt32, math.MaxInt32
		if idx-1 >= 0 {
			l = house - heaters[idx-1]
		}
		if idx < len(heaters) {
			r = heaters[idx] - house
		}
		ret = mymax(ret, mymin(l, r))
	}
	return ret
}
func mymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func mymin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
