package main

/**
dp算法：
令dp[n]为n对应的最大积。
那么递推方程就是：dp[n]=max(i*dp[n-i],i*(n-i))(其中i从1到n-1)。
边界:dp[2]=1;
*/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = -1
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i-j]*j, max(dp[i], j*(i-j)))
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
