package main

import "strings"

func reverseWords(s string) string {
	sp := strings.Split(s, " ")
	ret := ""
	for i := 0; i < len(sp); i++ {
		t := make([]byte, len(sp[i]))
		for left, right := 0, len(sp[i])-1; left <= right; left, right = left+1, right-1 {
			t[left], t[right] = sp[i][right], sp[i][left]
		}
		ret += string(t)
		if i != len(sp)-1 {
			ret += " "
		}
	}
	return ret

}
