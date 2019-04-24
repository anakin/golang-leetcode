package main

import "math"

/**
只能进行两次交易，因此一共有四个状态，互相之间转移
*/

func maxProfit(prices []int) int {
	b1, b2 := math.MinInt32, math.MinInt32
	s1, s2 := 0, 0
	for i := 0; i < len(prices); i++ {
		b1 = max(b1, -prices[i])
		s1 = max(s1, b1+prices[i])
		b2 = max(b2, s1-prices[i])
		s2 = max(s2, b2+prices[i])
	}
	return s2
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
