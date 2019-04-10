package main

/**
假设原数组为S，目标值为target，那么原数组必然可以分成两个部分，一个部分里面的元素前面需要加-，即运算的时候应该是做减法，另一个部分里面的元素前面需要加+，即运算的时候应该是做加法；
我们将做加法部分的数组记为P，做减法部分的数组记为N，举个例子，例如S = {1，2，3，4，5}，target = 3，那么有一种可以是1-2+3-4+5，即P = {1，3，5}，N = {2，4}；
于是我们可以知道：target = sum(P) - sum(N)；
那么sum(P) + sum(N) + sum(P) - sum(N) = sum(S) + target = 2sum(P)；
那么sum(P) = [target + sum(S)] / 2；
根据以上的推导，我们可以得到这样的结论：我们需要找到这样一个子序列P，使得子序列之和等于原序列之和与目标值的和的一半，我们需要找到这样子序列的数目。
显而易见，如果原序列之和与目标值的和为奇数，肯定不存在这样的子序列，如果原序列之和小于目标值，也肯定不存在（因为这种情况下，即使我们填的都是+，让所有元素相加都小于目标值，肯定不存在满足条件的解决方案）。
排除了以上的两种特殊情况之后，我们接下来研究的就是在原数组中找到之和符合条件的子序列个数，非常明显DP的方法。

*/
func findTargetSumWays(nums []int, S int) int {
	S = myabs(S)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < S || (S+sum)%2 != 0 {
		return 0
	}
	target := (S + sum) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}
func myabs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
