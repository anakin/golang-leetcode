package main

import "sort"

/**
dp
从最小的数开始
dp[c] = sum{dp[a]*dp[b]};c=a*b;a,b,c in A
ans = sum(dp[c])
*/
func numFactoredBinaryTrees(A []int) int {
	mod := 1000000007
	sort.Ints(A)
	dp := make(map[int]int)
	for i := 0; i < len(A); i++ {
		dp[A[i]] = 1
		for j := 0; j < i; j++ {
			if A[i]%A[j] == 0 {
				if _, ok := dp[A[i]/A[j]]; ok {
					dp[A[i]] += (dp[A[j]] * dp[A[i]/A[j]]) % mod
				}
			}
		}
	}
	ans := 0
	for _, v := range dp {
		ans += v
	}
	return ans % mod
}
