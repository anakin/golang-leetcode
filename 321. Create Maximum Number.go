package main

import (
	"strconv"
	"strings"
)

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	ans := []int{}
	if len(nums1) == 0 && len(nums2) == 0 {
		return ans
	}
	if len(nums1) == 0 {
		return nums2[:k]
	}
	if len(nums2) == 0 {
		return nums1[:k]
	}
	for i := 0; i < k; i++ {
		if i > len(nums1) || k-i > len(nums2) {
			continue
		}
		m := mergeArray(maxArray(nums1, i), maxArray(nums2, k-i))
		if arrayToStr(ans) < arrayToStr(m) {
			ans = m
		}
	}
	return ans
}

//数组转换成字符串，用于比较大小
func arrayToStr(nums []int) string {
	ret := ""
	str := []string{}
	if len(nums) == 0 {
		return ret
	}
	for _, v := range nums {
		str = append(str, strconv.Itoa(v))
	}
	return strings.Join(str, "")
}

//合并两个数组，不改变顺序的情况下，获得一个最大的数组
func mergeArray(nums1, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	ans := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if bigArray(nums1[i:], nums2[j:]) {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}
	if i < len(nums1) {
		ans = append(ans, nums1[i:]...)
	}
	if j < len(nums2) {
		ans = append(ans, nums2[j:]...)
	}
	return ans
}

//取数组中前k个最大的值
func maxArray(nums []int, k int) []int {
	ans := []int{}
	if k == 0 {
		return ans
	}
	l := len(nums)
	toPop := 0
	if l <= k {
		return nums
	} else {
		toPop = l - k
	}

	for _, v := range nums {
		for len(ans) > 0 && v > ans[len(ans)-1] && toPop > 0 {
			ans = ans[:len(ans)-1]
			toPop--
		}
		ans = append(ans, v)
	}
	return ans[:k]
}

//比较两个数组的大小，如果前一个大于后一个，返回true

func bigArray(nums1, nums2 []int) bool {
	l1, l2 := len(nums1), len(nums2)
	l := 0
	if l1 > l2 {
		l = l2
	} else {
		l = l1
	}
	for i := 0; i < l; i++ {
		if nums1[i] == nums2[i] {
			continue
		}
		if nums1[i] > nums2[i] {
			return true
		} else {
			return false
		}
	}
	if l1 > l2 {
		return true
	}
	return false
}
