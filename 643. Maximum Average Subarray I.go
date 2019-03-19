package main

/**
先计算前k个的sum，然后一直向右滑动，统计最大值
*/
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	m := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		if sum > m {
			m = sum
		}
	}
	return float64(m) / float64(k)
}
