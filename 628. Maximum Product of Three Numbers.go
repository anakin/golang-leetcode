package main

import (
	"math"
)

/**
考虑到有负数的情况，最大值有两种可能
1.三个最大的正数
2.两个最小的负数和一个最大的正数
*/

func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt32, math.MaxInt32
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	for _, v := range nums {
		if v <= min1 {
			min2, min1 = min1, v
		} else if v <= min2 {
			min2 = v
		}
		if v >= max1 {
			max3, max2, max1 = max2, max1, v
		} else if v >= max2 {
			max3, max2 = max2, v
		} else if v >= max3 {
			max3 = v
		}
	}
	m1 := min1 * min2 * max1
	m2 := max1 * max2 * max3
	if m1 > m2 {
		return m1
	}
	return m2
}
