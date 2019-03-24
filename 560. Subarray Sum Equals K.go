package main

/**
O(n^2)
*/
func subarraySum(nums []int, k int) int {
	l := len(nums)
	sums := make([]int, l+1)
	for i := 1; i <= l; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	ans := 0
	for i := 0; i < l; i++ {
		for j := i; i < l; j++ {
			if (sums[j+1] - sums[i]) == k {
				ans++
			}
		}
	}
	return ans
}

/**
prefix sum
O(n)
*/
func subarraySum1(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]int)
	m[0] = 1
	ans, sum := 0, 0
	for _, num := range nums {
		sum += num
		ans += m[sum-k]
		m[sum]++
	}
	return ans
}
