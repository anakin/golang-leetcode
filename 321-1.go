package main

/*
 * @lc app=leetcode id=321 lang=golang
 *
 * [321] Create Maximum Number
 *
 * https://leetcode.com/problems/create-maximum-number/description/
 *
 * algorithms
 * Hard (25.16%)
 * Total Accepted:    28.7K
 * Total Submissions: 114K
 * Testcase Example:  '[3,4,6,5]\n[9,1,2,5,8,3]\n5'
 *
 * Given two arrays of length m and n with digits 0-9 representing two numbers.
 * Create the maximum number of length k <= m + n from digits of the two. The
 * relative order of the digits from the same array must be preserved. Return
 * an array of the k digits.
 *
 * Note: You should try to optimize your time and space complexity.
 *
 * Example 1:
 *
 *
 * Input:
 * nums1 = [3, 4, 6, 5]
 * nums2 = [9, 1, 2, 5, 8, 3]
 * k = 5
 * Output:
 * [9, 8, 6, 5, 3]
 *
 * Example 2:
 *
 *
 * Input:
 * nums1 = [6, 7]
 * nums2 = [6, 0, 4]
 * k = 5
 * Output:
 * [6, 7, 6, 0, 4]
 *
 * Example 3:
 *
 *
 * Input:
 * nums1 = [3, 9]
 * nums2 = [8, 9]
 * k = 3
 * Output:
 * [9, 8, 9]
 *
 */

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 如果nums1>nums2，返回true；否则返回false
func compare(nums1, nums2 []int) bool {
	i := 0
	for ; i < len(nums1) && i < len(nums2); i++ {
		if nums1[i] > nums2[i] {
			return true
		} else if nums2[i] > nums1[i] {
			return false
		}
	}
	if i == len(nums1) {
		return false
	}
	return true
}

/*
思路： 2*greedy + 1*DP

子问题1：找到数组中由k个数组成的最大数的组合（不改变先后顺序）
	getMax(nums []int, k int) []int {}

子问题2：将长度为i和k-i的两个最大值组合进行合并(不改变先后顺序)
	mergeMax(max1, max2 []int) []int {}

子问题3：获取子问题2中所有合并所得结果的最大组合
	dpMax(max1, max2 []int) []int {}
*/
func getMax(nums []int, k int) []int {
	size := len(nums)
	res := make([]int, k)
	// j代表res的长度，不是index
	j := 0
	for i := 0; i < size; i++ {
		for j > 0 && nums[i] > res[j-1] && size-i > k-j {
			j--
		}
		if j < k {
			res[j] = nums[i]
			j++
		}
	}
	return res
}

func mergeMax(max1, max2 []int, k int) []int {
	res := make([]int, k)
	i1, i2 := 0, 0
	n1, n2 := len(max1), len(max2)
	for i := 0; i < k; i++ {
		if i1 == n1 {
			res[i] = max2[i2]
			i2++
			continue
		}
		if i2 == n2 {
			res[i] = max1[i1]
			i1++
			continue
		}
		//if max1[i1] > max2[i2] {
		if compare(max1[i1:], max2[i2:]) {
			res[i] = max1[i1]
			i1++
		} else {
			res[i] = max2[i2]
			i2++
		}
	}
	return res
}

func dpMax(max1, max2 []int, k int) []int {
	for i := 0; i < k; i++ {
		if max1[i] > max2[i] {
			return max1
		} else if max2[i] > max1[i] {
			return max2
		}
	}
	return max1
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n2 < n1 {
		return maxNumber(nums2, nums1, k)
	}
	iMax := min(n1, k)
	res := make([]int, k)
	for i := 0; i <= iMax; i++ {
		if k-i <= n2 {
			res = dpMax(res, mergeMax(getMax(nums1, i), getMax(nums2, k-i), k), k)
		}
	}
	return res
}
