package main

/**
双重遍历，可以排除10的倍数
*/
func selfDividingNumbers(left int, right int) []int {
	ret := []int{}
	for i := left; i <= right; i++ {
		flag := true
		for tmp := i; flag && tmp > 0; {
			if tmp%10 == 0 || i%(tmp%10) > 0 {
				flag = false
			}
			tmp /= 10
		}
		if flag {
			ret = append(ret, i)
		}
	}
	return ret
}
