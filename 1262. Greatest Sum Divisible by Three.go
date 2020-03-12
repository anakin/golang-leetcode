/**
Given an array nums of integers, we need to find the maximum possible sum of elements of the array such that it is divisible by three.



Example 1:

Input: nums = [3,6,5,1,8]
Output: 18
Explanation: Pick numbers 3, 6, 1 and 8 their sum is 18 (maximum sum divisible by 3).
Example 2:

Input: nums = [4]
Output: 0
Explanation: Since 4 is not divisible by 3, do not pick any number.

*/
package main

func maxSumDivThree(nums []int) int {
	dp := make([]int, 3)
	for _, v := range nums {
		tmp := make([]int, 3)
		copy(tmp, dp)
		for i := 0; i < 3; i++ {
			dp[(v+tmp[i])%3] = max(dp[(v+tmp[i])%3], v+tmp[i])
		}
	}
	return dp[0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/**
dp:除以3的余数有3种情况，0，1，2
*/
