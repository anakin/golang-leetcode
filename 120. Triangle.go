package main

/**
dp算法
dp数组保存当前行获得的最小值
从下往上转移
初始化为最后一行的值
*/

func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return triangle[0][0]
	}
	dp := make([]int, l)
	//dp数组初始化成最后一行的值
	for i := 0; i < l; i++ {
		dp[i] = triangle[l-1][i]
	}
	for i := l - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x

}
