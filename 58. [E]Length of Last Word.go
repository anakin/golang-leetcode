package main

import (
	"strings"
)

/**
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5
*/
func lengthOfLastWord(s string) int {
	arr := strings.Split(strings.Trim(s, " "), " ")
	length := len(arr)
	res := 0
	for i := length - 1; i >= 0; i-- {
		if arr[i] == " " {
			continue
		} else {
			res = len(arr[i])
			break
		}
	}
	return res
}
