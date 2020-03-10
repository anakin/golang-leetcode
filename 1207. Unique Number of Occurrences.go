/**
Given an array of integers arr, write a function that returns true if and only if the number of occurrences of each value in the array is unique.



Example 1:

Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
Example 2:

Input: arr = [1,2]
Output: false

*/
package main

func uniqueOccurrences(arr []int) bool {
	counts := make(map[int]int)
	for _, i := range arr {
		counts[i] += 1
	}
	set := make(map[int]bool)
	for _, c := range counts {
		if _, ok := set[c]; ok {
			return false
		}
		set[c] = true
	}
	return true
}
