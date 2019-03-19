package main

import "math"

//从0到sqrt（c）之间进行遍历，检测剩下的数开方后是不是整数
func judgeSquareSum(c int) bool {
	limit := int(math.Sqrt(float64(c)))
	for i := 0; i <= limit; i++ {
		k := c - i*i
		tmp := int(math.Sqrt(float64(k)))
		if tmp*tmp == k {
			return true
		}

	}
	return false
}
