package main

import (
	"math"
)

/**
用两个变量记录出现最多的两个元素(大于1/3的元素不可能出现3个）
然后统计出现这两个元素的个数,符合条件就加入返回值中
*/
func majorityElement(nums []int) []int {
	l := len(nums)
	if l == 0 {
		return []int{}
	}
	if l == 1 {
		return nums
	}
	m1, m2 := nums[0], math.MaxInt32
	c1, c2 := 1, 0
	for i := 1; i < l; i++ {
		if nums[i] == m1 {
			c1++
		} else if nums[i] == m2 {
			c2++
		} else if c1 == 0 {
			m1 = nums[i]
			c1 = 1
		} else if c2 == 0 {
			m2 = nums[i]
			c2 = 1
		} else {
			c1--
			c2--
		}
	}
	t1, t2 := 0, 0
	for _, v := range nums {
		if v == m1 {
			t1++
		}
		if v == m2 {
			t2++
		}
	}
	ret := []int{}
	if t1 > l/3 {
		ret = append(ret, m1)
	}
	if t2 > l/3 {
		ret = append(ret, m2)
	}
	return ret
}
