package main

/**
用hash表判断重复项
用累加和减去不重复项之和，等于丢失的数字
*/
//TODO 速度待优化
func findErrorNums(nums []int) []int {

	m := make(map[int]int)
	sum := 0
	ret := make([]int, 2)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			ret[0] = v
		} else {
			m[v] = 1
			sum += v
		}
	}
	l := len(nums)
	ret[1] = l*(1+l)/2 - sum
	return ret
}

/**
用一个新数组保存出现的次数，然后找出出现两次的和出现0次的
速度比上面快很多
*/
func findErrorNums1(nums []int) []int {
	l := len(nums)
	m := make([]int, l)
	ret := make([]int, 2)
	for i := 0; i < l; i++ {
		m[nums[i]-1]++
	}
	for k, v := range m {
		if v == 2 {
			ret[0] = k + 1
		}
		if v == 0 {
			ret[1] = k + 1
		}
	}
	return ret
}
