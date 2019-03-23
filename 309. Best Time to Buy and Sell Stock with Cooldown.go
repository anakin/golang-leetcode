package main

import "math"

/**
三种状态之间互相转移：持有，售出，休息
售出的收益等于持有加上当天价格
持有的收益等于max(前一天休息之后买入，前一天持有）
休息的收益等于max(前一天休息，前一天售出)
*/
func maxProfit(prices []int) int {
	sold, rest, hold := 0, 0, math.MinInt32
	for _, v := range prices {
		pre_sold := sold
		sold = hold + v
		hold = mymax(hold, rest-v)
		rest = mymax(rest, pre_sold)
	}
	return mymax(rest, sold)
}

func mymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
