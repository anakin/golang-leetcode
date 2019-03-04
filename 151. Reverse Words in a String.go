package main

import (
	"fmt"
	"strings"
)

//总觉得有更好的办法
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	res := []string{}
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			res = append(res, arr[i])
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return strings.Replace(strings.Trim(fmt.Sprint(res), "[]"), " ", " ", -1)

}
