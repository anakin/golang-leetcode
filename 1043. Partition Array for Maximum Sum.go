/**
Given an integer array A, you partition the array into (contiguous) subarrays of length at most K.  After partitioning, each subarray has their values changed to become the maximum value of that subarray.

Return the largest sum of the given array after partitioning.



Example 1:

Input: A = [1,15,7,9,2,5,10], K = 3
Output: 84
Explanation: A becomes [15,15,15,9,10,10,10]

*/

package main

func maxSumAfterPartitioning(A []int, K int) int {
	n := len(A)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		curMax := A[i]
		for j := 1; j <= K && i-j+1 >= 0; j++ {
			curMax = max(curMax, A[i-j+1])
			dp[i+1] = max(dp[i+1], dp[i-j+1]+curMax*j)
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
note:DP
*/
