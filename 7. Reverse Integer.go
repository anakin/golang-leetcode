package main

import "math"

/**
循环取余数
需要注意越界
*/
func reverse(x int) int {

	res := 0
	for x != 0 {
		carry := x % 10
		res = res*10 + carry
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return res
}
