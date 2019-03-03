package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x != 0 && x%10 == 0 {
		return false
	}
	y := 0
	c := x
	for x != 0 {
		y = y*10 + x%10
		x /= 10
	}
	return c == y
}
