/**
Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all elements in arr2 are also in arr1.

Sort the elements of arr1 such that the relative ordering of items in arr1 are the same as in arr2.  Elements that don't appear in arr2 should be placed at the end of arr1 in ascending order.



Example 1:

Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
Output: [2,2,2,1,4,3,3,9,6,7,19]

*/
package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	counter := map[int]int{}
	others := []int{}
	dest := make([]int, 0, len(arr2))
	for _, v := range arr2 {
		counter[v] = 0
	}
	for _, v := range arr1 {
		if _, ok := counter[v]; ok {
			counter[v]++
		} else {
			others = append(others, v)
		}
	}

	for _, v := range arr2 {
		for i := 0; i < counter[v]; i++ {
			dest = append(dest, v)
		}
	}
	sort.Ints(others)
	dest = append(dest, others...)
	return dest
}
