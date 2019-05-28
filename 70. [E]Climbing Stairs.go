package main

/**
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

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
