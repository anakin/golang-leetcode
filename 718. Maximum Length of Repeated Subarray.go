package main

/**
DP算法

dp[i][j]表示A和B共同前缀的长度
初始都是0
转移方程：dp[i][j] = dp[i+1][j+1]+1
*/
func findLength(A []int, B []int) int {
	ans := 0
	la, lb := len(A), len(B)
	dp := make([][]int, la+1)
	for k := range dp {
		dp[k] = make([]int, lb+1)
	}
	for i := la - 1; i >= 0; i-- {
		for j := lb - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if ans < dp[i][j] {
					ans = dp[i][j]
				}
			}
		}
	}
	return ans
}
