package main

/**
回溯算法搞定
*/

var res [][]string

func partition(s string) [][]string {
	res = [][]string{}
	if len(s) < 1 {
		return res
	}
	helper(s, 0, []string{})
	return res
}
func helper(s string, start int, tmp []string) {
	if start == len(s) {
		m := make([]string, len(tmp))
		copy(m, tmp)
		res = append(res, m)
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			tmp = append(tmp, s[start:i+1])
			helper(s, i+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
