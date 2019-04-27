package main

/**
把nums复制两倍，保证从任何一个位置开始，后面包含完整的数据
*/

func nextGreaterElements(nums []int) []int {
	l := len(nums)
	l2 := 2 * l
	ret := make([]int, l)
	index := make([]int, l2)
	for i := l; i < l2; i++ {
		index[i] = i - l
	}
	cur := l
	for i := l - 1; i >= 0; i-- {
		for cur < l2 && nums[index[cur]] <= nums[i] {
			cur++
		}
		if cur == l2 {
			ret[i] = -1
		} else {
			ret[i] = nums[index[cur]]
		}
		cur--
		index[cur] = i
	}
	return ret
}
