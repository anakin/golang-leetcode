package main

import (
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return count(countAndSay(n - 1))
}

func count(s string) string {
	c := string(s[0])
	count := 1
	res := ""
	for _, char := range s[1:] {
		if string(char) == c {
			count += 1
		} else {
			res = res + strconv.Itoa(count) + string(c)
			c = string(char)
			count = 1
		}
	}
	res = res + strconv.Itoa(count) + string(c)
	return res
}
