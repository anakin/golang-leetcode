package main

/**
dp
用bitmap保存中间状态
*/
func helper(maxInt, left, bit int, dp map[int]bool, bitmap []int) bool {
	if val, ok := dp[bit]; ok {
		return val
	}
	if left <= maxInt {
		for i := maxInt; i >= left; i-- {
			if (bit & bitmap[i]) == 0 {
				dp[bit] = true
				return true
			}
		}
	}
	for i := maxInt; i > 0; i-- {
		if (bit & bitmap[i]) > 0 {
			continue
		}
		bit |= bitmap[i]
		ret := helper(maxInt, left-i, bit, dp, bitmap)
		bit &= ^bitmap[i]
		if ret == false {
			dp[bit] = true
			return true
		}
	}
	dp[bit] = false
	return false
}
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	dp := make(map[int]bool)
	sum := 0
	bitmap := make([]int, maxChoosableInteger+1)
	for i := maxChoosableInteger; i > 0; i-- {
		sum += i
		bitmap[i] = 1 << uint8(i)
	}
	if sum < desiredTotal {
		return false
	}
	return helper(maxChoosableInteger, desiredTotal, 0, dp, bitmap)

}
