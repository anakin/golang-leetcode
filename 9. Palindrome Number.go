package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	//排除以0结尾的情况
	if x != 0 && x%10 == 0 {
		return false
	}
	y := 0
	c := x
	//把x反转，判断是否和原来的x相等
	for x != 0 {
		y = y*10 + x%10
		x /= 10
	}
	return c == y
}
