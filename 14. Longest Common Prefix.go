package main

import "math"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := math.MaxInt32
	minStr := ""
	//先找到最短的字符串
	for _, v := range strs {
		l := len(v)
		if l < minLen {
			minLen = l
			minStr = v
		}
	}
	long := len(minStr)
	for _, v := range strs {
		for i := 0; i < minLen; i++ {
			if minStr[i:i+1] != v[i:i+1] {
				long = int(math.Min(float64(long), float64(i)))
			}
		}
	}
	return minStr[:long]
}
