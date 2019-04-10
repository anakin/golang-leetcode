package main

/**
如果num有约数d，则它必有约数num/d
看到变态的解法：
根据题目给的范围可知，范围内的完美数只有：6, 28, 496, 8128, 33550336，枚举就行了
*/
func checkPerfectNumber(num int) bool {
	if num <= 0 {
		return false
	}
	sum := 0
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i*i != num {
				sum += num / i
			}
		}
	}
	return sum-num == num
}
