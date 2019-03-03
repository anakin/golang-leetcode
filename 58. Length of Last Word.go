package main

import (
	"strings"
)

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
