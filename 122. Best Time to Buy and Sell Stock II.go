package main

import "math"

func maxProfit(prices []int) int {
	cur, max := 0, 0
	for i := 1; i < len(prices); i++ {
		cur = mymax(cur, cur+prices[i]-prices[i-1])
		max = mymax(cur, max)
	}
	return max
}
func mymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
