package main

/**
greedy算法
从两边向中间搜索，如果对应位置不相等
则去掉左边的或者去掉右边的元素之后，剩余部分应该为回文
*/
func validPalindrome(s string) bool {
	str := []byte(s)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-i-1] {
			j := l - i - 1
			return isPalindrome(str, i+1, j) || isPalindrome(str, i, j-1)
		}
	}
	return true
}
func isPalindrome(s []byte, l, r int) bool {
	for i := l; i <= l+(r-l)/2; i++ {
		if s[i] != s[r-i+l] {
			return false
		}
	}
	return true
}
