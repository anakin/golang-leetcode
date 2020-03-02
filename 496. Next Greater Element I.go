package main

/**
nums1 是nums2的子集，因此find方法可以找到nums1中的每个元素在nums2中的位置
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ret := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		index := find(nums2, nums1[i])
		for j := index; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				ret[i] = nums2[j]
				break
			}
		}
		if ret[i] == 0 {
			ret[i] = -1
		}
	}
	return ret
}

func find(A []int, a int) int {
	for i := 0; i < len(A); i++ {
		if A[i] == a {
			return i
		}
	}
	return -1
}

/*
用栈来解决，如果栈顶的元素比当前的值大，那么栈顶元素就是要找的值，记录当前元素和栈顶元素的对应关系并出栈
*/
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	data := make(map[int]int, len(nums2))
	var stack []int
	for _, num := range nums2 {
		for len(stack) != 0 && stack[len(stack)-1] < num {
			data[stack[len(stack)]-1] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	res := make([]int, len(nums1))
	for i, num := range nums1 {
		if elem, ok := data[num]; ok {
			res[i] = elem
		} else {
			res[i] = -1
		}
	}
	return res
}
