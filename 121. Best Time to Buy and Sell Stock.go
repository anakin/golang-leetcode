package main

import "math"

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := math.MaxInt64
	for _, v := range prices {
		minPrice = mymin(minPrice, v)
		maxProfit = mymax(maxProfit, v-minPrice)
	}
	return maxProfit
}
func mymin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func mymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
