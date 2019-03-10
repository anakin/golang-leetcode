package main

func longestPalindrome(s string) int {
	count := make([]int, 128)
	for _, c := range s {
		count[c]++
	}
	ans := 0
	for _, v := range count {
		ans += v / 2 * 2
		if ans%2 == 0 && v%2 == 1 {
			ans++
		}
	}
	return ans
}
