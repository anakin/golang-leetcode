package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := getHeight(root.Left)
	r := getHeight(root.Right)
	return l-r < 2 && r-l < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return int(math.Max(float64(getHeight(root.Left)), float64(getHeight(root.Right))) + 1)
}
