package main

/**
dp解法
状态方程：dp[i] =min(dp[i-2]+cost[i-2],dp[i-1]+cost[i-1])
*/

func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	t := make([]int, l)
	t[0] = cost[0]
	t[1] = cost[1]
	for i := 2; i < l; i++ {
		t[i] = mymin(t[i-1]+cost[i], t[i-2]+cost[i])
	}
	return mymin(t[l-1], t[l-2])
}

func mymin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
