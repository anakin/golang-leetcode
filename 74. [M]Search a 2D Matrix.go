package main

/**
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
Example 1:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
Output: true
Example 2:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
Output: false
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	n := len(matrix)
	m := len(matrix[0])
	if target < matrix[0][0] || target > matrix[n-1][m-1] {
		return false
	}

	var mid int
	low, high := 0, n-1
	for low <= high {
		mid = (low + high) / 2
		if matrix[mid][0] > target {
			high = mid - 1
		} else if matrix[mid][0] < target {
			low = mid + 1
		} else {
			return true
		}
	}

	left, right := 0, m-1
	for left <= right {
		mid = (left + right) / 2
		if matrix[high][mid] > target {
			right = mid - 1
		} else if matrix[high][mid] < target {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}
