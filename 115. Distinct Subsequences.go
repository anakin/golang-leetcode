package main

/**
dp
dp[i][j]表示构成i长度的t，用到j长度的s，结果等于种类
转移方程：
如果t[i]==s[j],dp[i][j]=dp[i-1][j-1]+dp[i][j-1]
否则：dp[i][j]=dp[i][j-1]
初始值：
dp[0][*]=1表示构成0长度的t，有一种方法
*/
func numDistinct(s string, t string) int {
	ls, lt := len(s), len(t)
	dp := make([][]int, lt+1)
	for k := range dp {
		dp[k] = make([]int, ls+1)
	}
	for k := range dp[0] {
		dp[0][k] = 1
	}
	for i := 1; i <= lt; i++ {
		for j := 1; j <= ls; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[lt][ls]
}
