package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if v <= i {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]
}
