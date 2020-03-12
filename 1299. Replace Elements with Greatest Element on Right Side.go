/**
Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.

After doing so, return the array.



Example 1:

Input: arr = [17,18,5,4,6,1]
Output: [18,6,6,6,1,-1]

*/
package main

func replaceElements(arr []int) []int {
	var res []int = make([]int, len(arr))
	var max int = arr[len(arr)-1]
	res[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		res[i] = max
		if arr[i] > max {
			max = arr[i]
		}
	}
	return res
}

/**
trick:from right to left
*/
