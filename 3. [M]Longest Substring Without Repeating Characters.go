package main

/**
Given a string, find the length of the longest substring without repeating characters.
Example 1:
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

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
