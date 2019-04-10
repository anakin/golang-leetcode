package main

/**
遍历整个数组，依次加当前数组元素并将相加和求余k，求余结果只有0~k-1这k中情况，将求余结果存入Hash Table中。如果遍历到当前位置求余结果已经在Hash Table中，表明从上一求余结果相同的位置到当前位置的子数组相加和是k的倍数，否则将求余结果存入Hash Table
*/
func checkSubarraySum(nums []int, k int) bool {
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}
		prev, ok := m[sum]
		if ok {
			if i-prev > 1 {
				return true
			}
		} else {
			m[sum] = i
		}
	}
	return false
}
