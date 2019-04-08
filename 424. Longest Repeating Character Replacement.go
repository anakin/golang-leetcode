package main

/**
定义start和end两个标记，中间的内容即是窗口，计算窗口内所有字母出现的次数，因为全是大写字母，所以可以用一个26位的数组来记录窗口内每个字母出现的次数。
找到窗口内出现最多的次数，加上允许替换的字母数k，看是否超过窗口宽度，如果超过了，说明窗口还可以更长， 也就是说窗口内重复的字母的长度可以更长，就将end右移一位，形成新的窗口，然后继续重复上面的步骤。如果没超过，说明能构成的最长的重复字母长度已经到顶了，这时应该将start右移一位，来寻找新的可能的更长重复字母长度。
每次计算重复字母长度时，当出现更长的可能时，都更新最终的结果。
*/

func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	str := []byte(s)
	start, end := 0, 0
	ret := 0
	c := make([]int, 26)
	c[str[0]-'A']++
	for len(str) > end {
		maxc := 0
		for i := 0; i < 26; i++ {
			if c[i] > maxc {
				maxc = c[i]
			}
		}
		if maxc+k > end-start {
			end++
			if end < len(str) {
				c[str[end]-'A']++
			}
		} else {
			c[str[start]-'A']--
			start++
		}
		if maxc+k > ret {
			if maxc+k <= len(str) {
				ret = maxc + k
			} else {
				ret = len(str)
			}
		}
	}
	return ret
}
