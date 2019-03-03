package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	start, max := -1, 0

	for k, v := range s {
		if last, ok := m[v]; ok && last > start {
			start = last
		}
		m[v] = k
		if k-start > max {
			max = k - start
		}
	}
	return max
}
