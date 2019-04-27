package main

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for k, _ := range dp {
		dp[k] = make([]int, n+1)
	}
	for _, v := range strs {
		c0, c1 := 0, 0
		for i := 0; i < len(v); i++ {
			if v[i] == '1' {
				c1++
			} else {
				c0++
			}
		}
		if c0 > m || c1 > n {
			continue
		}
		for i := m; i >= c0; i-- {
			for j := n; j >= c1; j-- {
				dp[i][j] = max(dp[i][j], dp[i-c0][j-c1]+1)
			}
		}
	}
	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
