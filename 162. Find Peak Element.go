package main

//不难，就是要注意首尾的判断
func findPeakElement(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	if nums[0] > nums[1] { //第一个元素
		return 0
	} else if nums[l-1] > nums[l-2] { //最后一个元素
		return l - 1
	} else { //中间的
		for i := 1; i < len(nums)-1; i++ {
			if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
				return i
			}
		}
	}
	return -1
}
