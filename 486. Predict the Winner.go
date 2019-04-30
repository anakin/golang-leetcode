package main

/**
DP
dp[i][j]表示原数组中从i到j的这么多数中，按照游戏规则，某个玩家所能获得的最大分数。
假设这个分数此时属于palyer1，那么dp[i+1][j]或者dp[i][j-1]表示player2玩家所能获得的最大分数。因为对于player1来讲，他第一次选择要么是第i个数，要么是第j个数，所以对于player2来讲，就分两种情况取最大。
另外我们设从i到j的所有数的和是sum[i,j]，则可以得到递推公式：（动态规划最明显的标识）
dp[i][j]=max(sum[i+1][j]-dp[i+1][j]+nums[i], sum[i][j-1]-dp[i][j-1]+nums[j]) 。
化简一下：
dp[i][j]=max(sum[i][j]-dp[i+1][j], sum[i][j]-dp[i][j-1]) 。
*/
func PredictTheWinner(nums []int) bool {
	l := len(nums)
	dp := make([][]int, l+1)
	for k := range dp {
		dp[k] = make([]int, l)
	}
	for i := l; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			a := nums[i] - dp[i+1][j]
			b := nums[j] - dp[i][j-1]
			dp[i][j] = max(a, b)
		}
	}
	return dp[0][l-1] >= 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
