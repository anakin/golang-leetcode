package main

func isHappy(n int) bool {
	return isOne(n)
}
func getSum(n int) int {
	if n/10 == 0 {
		return n * n
	}
	return (n%10)*(n%10) + getSum(n/10)
}
func isOne(n int) bool {
	a := getSum(n)
	if a == 1 {
		return true
	} else if a == 4 {
		return false
	} else {
		return isOne(a)
	}
}
