package main

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	pos := 1
	res := 0
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return res
	}
	i := 0
	if str[i] == '+' {
		i++
		pos = 1
	} else if str[i] == '-' {
		i++
		pos = -1
	}
	for ; i < len(str); i++ {
		if pos*res >= math.MaxInt32 {
			return math.MaxInt32
		}
		if pos*res <= math.MinInt32 {
			return math.MinInt32
		}
		if str[i] < '0' || string(str[i]) > "9" {
			return res * pos
		}
		res = res*10 + int(str[i]) - '0'
	}
	if pos*res >= math.MaxInt32 {
		return math.MaxInt32
	}
	if pos*res <= math.MinInt32 {
		return math.MinInt32
	}
	return pos * res
}
