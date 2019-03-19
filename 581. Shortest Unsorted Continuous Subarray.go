package main

/**
来看这个例子：1 3 5 7 2 4 5 6

　　a. 当我们找到第一个违反ascending 排序的数字 2的时候，我们不能是仅仅把beg 标记为2的前面一个数字7，而是要一直往前，找到一个合适的位置，找到在最前面位置的比2大的数字，这里是3。

　　b. 同样的，为了找end， 那么我们要从7的后面开始找，一直找到一个最后面位置的比7小的数字，这里是6。
*/
func findUnsortedSubarray(nums []int) int {

	n := len(nums)
	begin, end := -1, -2
	min, max := nums[n-1], nums[0]
	for i := 0; i < n; i++ {
		max = mymax(max, nums[i])
		min = mymin(min, nums[n-1-i])
		if nums[i] < max {
			end = i
		}
		if nums[n-1-i] > min {
			begin = n - 1 - i
		}
	}
	return end - begin + 1
}

//使用自定义的max、min，速度会快很多
func mymin(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func mymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
