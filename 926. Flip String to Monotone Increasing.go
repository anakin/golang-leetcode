package main

/**
dp解法
状态方程：
如果最后一位是0：
	dp[i][0] = dp[i-1][0]
	dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
如果最后一位是1：
	dp[i][0] = dp[i-1][0] + 1
	dp[i][1] = min(dp[i-1][0], dp[i-1][1])

TODO 可以降维优化
*/

func minFlipsMonoIncr(S string) int {
	l := len(S)
	dp := make([][]int, l+1)
	//二维数组初始化方法
	for i := 0; i < l+1; i++ {
		dp[i] = make([]int, 2)
	}
	for i := 1; i <= l; i++ {
		if S[i-1] == '0' {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = mymin(dp[i-1][0], dp[i-1][1]) + 1
		} else {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = mymin(dp[i-1][0], dp[i-1][1])
		}
	}
	return mymin(dp[l][0], dp[l][1])
}
func mymin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
