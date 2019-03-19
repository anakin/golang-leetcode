package main

/**
滑动窗口，如果在hash表里找到了当前的字符，重新设置窗口的起点位置
*/
func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	start, max := -1, 0

	for k, v := range s {
		if last, ok := m[v]; ok && last > start { //如果发现重复字符
			start = last
		}
		m[v] = k
		if k-start > max { //保存最大值
			max = k - start
		}
	}
	return max
}
