package main

import (
	"fmt"
	"regexp"
	"strings"
)

//TODO 速度不是很理想
func isPalindrome1(s string) bool {
	pat := "[,:.@#--?\";!` ]"
	re, _ := regexp.Compile(pat)

	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	fmt.Println("s=", s)
	if s == "" {
		return true
	}
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}
