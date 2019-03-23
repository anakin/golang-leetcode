package main

func findTargetSumWays(nums []int, S int) int {
	S = myabs(S)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < S || (S+sum)%2 != 0 {
		return 0
	}
	target := (S + sum) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}
func myabs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
