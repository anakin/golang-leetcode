package main

/**
分成两种情况考虑:
打劫第一个house
不打劫第一个house

如果打劫第一个house，就不能打劫最后一个house
否则就可以打劫最后一个
最后取这两种情况的最大值
*/
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	//dp1表示从第一个house开始
	//dp2表示从第二个house开始
	dp1, dp2 := make([]int, l), make([]int, l)
	dp1[0] = nums[0]
	dp1[1] = max(dp1[0], nums[1])
	dp2[0] = 0
	dp2[1] = max(dp2[0], nums[1])
	for i := 2; i < l; i++ {
		if i < l-1 {
			dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
		} else { //最后一个house
			dp1[i] = dp1[i-1] //不能抢
		}
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i])
	}
	return max(dp1[l-1], dp2[l-1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
