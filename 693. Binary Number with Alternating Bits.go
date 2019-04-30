package main

func hasAlternatingBits(n int) bool {
	cur := n % 2
	n /= 2
	for n > 0 {
		if cur == n%2 {
			return false
		}
		cur = n % 2
		n /= 2
	}
	return true
}
