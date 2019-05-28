package main

/**
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
Example 1:
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:
Input: "cbbd"
Output: "bb"
*/

/**
dp算法
*/

func longestPalindrome(s string) string {
	str := []byte(s)
	l := len(str)
	if l == 0 {
		return ""
	}
	dp := make([][]bool, l)
	for k := range dp {
		dp[k] = make([]bool, l)
	}
	ret := []byte{}
	max := 0
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if str[i] == str[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if max < j-i+1 {
					max = j - i + 1
					ret = str[i : j+1]
				}
			}
		}
	}
	return string(ret)
}
