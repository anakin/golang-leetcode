package main

func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	flags := make([]bool, l+1)
	flags[0] = true
	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			if flags[j] == true && contain(s[j:i], wordDict) {
				flags[i] = true
				break
			}
		}
	}
	return flags[l]
}

func contain(s string, wordDict []string) bool {
	for _, v := range wordDict {
		if s == v {
			return true
		}
	}
	return false
}
