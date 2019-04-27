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
