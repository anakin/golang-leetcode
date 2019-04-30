package main

import "strings"

func repeatedStringMatch(A string, B string) int {
	if len(A) == 0 && len(B) == 0 {
		return 1
	} else if len(A) == 0 || len(B) == 0 {
		return 0
	}
	s := A
	i := 0
	m := len(B)/len(A) + 2
	for i < m {
		if strings.Index(s, B) > -1 {
			return i + 1
		}
		s += A
		i++
	}
	return -1
}
