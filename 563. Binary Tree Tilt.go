package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum = 0

func findTilt(root *TreeNode) int {
	helper(root)
	return sum
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	right := helper(root.Right)
	sum += int(math.Abs(float64(left) - float64(right)))
	return root.Val + left + right
}
