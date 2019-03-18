package main

import "strings"

func checkRecord(s string) bool {
	t := strings.Count(s, "A")
	if t >= 2 {
		return false
	}
	t = strings.Count(s, "LLL")
	if t > 0 {
		return false
	}
	return true
}
