package main

import "math"

func maxProfit(prices []int) int {
	cur, max := 0, 0
	for i := 1; i < len(prices); i++ {
		cur = int(math.Max(float64(cur), float64(cur+prices[i]-prices[i-1])))
		max = int(math.Max(float64(cur), float64(max)))
	}
	return max
}
