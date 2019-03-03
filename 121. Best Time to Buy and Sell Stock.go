package main

import "math"

func maxProfit(prices []int) int {
	max_profit := 0
	min_price := math.MaxInt64
	for _, v := range prices {
		min_price = int(math.Min(float64(min_price), float64(v)))
		max_profit = int(math.Max(float64(max_profit), float64(v-min_price)))
	}
	return max_profit
}
