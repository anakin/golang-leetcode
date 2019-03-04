package main

//DP
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	tmp := []int{1, 2}
	for i := 2; i < n; i++ {
		tmp = append(tmp, tmp[i-1]+tmp[i-2])
	}
	return tmp[n-1]
}
