package main

/**
dp
dp[i]:是否可以得到和i
dp[i+v]=true,如果dp[i]==true，如果可以得到i，那么加上v就得到了i+v，所以dp[i+v]=true
dp[0] =true,得到和0是可以的
循环处理获得dp，期间要判断是否已经得到了结果
*/
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	dp := make([]bool, sum+1)
	for k := range dp {
		dp[k] = false
	}
	dp[0] = true
	for _, v := range nums {
		//注意，这里是倒着循环，避免被覆盖
		for i := sum; i >= 0; i-- {
			if dp[i] {
				dp[i+v] = true
			}
		}
		if dp[sum/2] {
			return true
		}
	}
	return false
}
