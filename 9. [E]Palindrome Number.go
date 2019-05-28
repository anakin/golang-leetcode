package main

/**
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

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
