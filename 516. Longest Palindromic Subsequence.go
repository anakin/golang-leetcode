package main

/**
dp:方程表示从i到j的长度中，回文子串的长度
Base case:
a ->dp[i][i]=1

case 1:s[i]==s[j]
	a*****a ->dp[i][j]=dp[i+1][j-1]+2

case 2:s[i]!=s[j]
	ab****b dp[i][j]=dp[i+1][i]
	a****ab dp[i][j]=dp[i][j-1]

*/
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for k := range dp {
		dp[k] = make([]int, n)
	}
	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			if i == j {
				dp[i][j] = 1
				continue
			}
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = mymax(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func mymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
