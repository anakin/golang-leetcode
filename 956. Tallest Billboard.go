package main

/**
dp思路
dp方程的键为两个柱子之间的高度差，值为当前高度差情况下，两个柱子的最小高度
状态转移的时候有三种情况，其中后两种可以合并
最后dp[0]保存的就是两个柱子高度差为0的时候，两个柱子的最小高度
*/
func tallestBillboard(rods []int) int {
	sum := 0
	for _, v := range rods {
		sum += v
	}
	dp := make([]int, sum+1)
	for k, _ := range dp {
		dp[k] = -1
	}
	dp[0] = 0
	for _, v := range rods {
		cur := make([]int, sum+1)
		copy(cur, dp)
		//for i := 0; i <= sum; i++ {
		for d, val := range cur {
			if val != -1 {
				if d+v <= sum {
					dp[d+v] = mymax(dp[d+v], val)
				}
				dp[myabs(d-v)] = mymax(dp[myabs(d-v)], val+mymin(v, d))
			}
		}
		//fmt.Println(dp)
	}
	return dp[0]
}
func mymin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func myabs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func mymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
