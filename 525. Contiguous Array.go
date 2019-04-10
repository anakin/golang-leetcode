package main

/**
把所有为0的数字都替换成-1，定义一个变量sum = 0，然后我们让这个变量与数组中的每个元素逐个相加，用map记录每次相加的结果及其下标，如果在后面得到某个结果(记该结果的下标为b)是map里面已经存在的值（记该值的下标为a），那么证明sum在a~b这个范围内相加的结果为0，也就是这个范围中0与1的数量相同，然后通过下标相减得到这个范围长度。接下来我们不断更新这种范围长度的最大值便可。
*/
func findMaxLength(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	ret, sum := 0, 0
	m := make(map[int]int)
	m[0] = -1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := m[sum]; ok {
			ret = max(ret, i-v)
		} else {
			m[sum] = i
		}
	}
	return ret
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
