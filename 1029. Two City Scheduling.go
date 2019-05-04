package main

import "sort"

/**
先排序
*/

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})
	sum := 0
	for i := 0; i < len(costs); i++ {
		if i < len(costs)/2 {
			sum += costs[i][0]
		} else {
			sum += costs[i][1]
		}
	}
	return sum
}
