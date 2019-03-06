package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root != nil {
		if isLeaf(root.Left) {
			sum += root.Left.Val
		} else {
			sum += sumOfLeftLeaves(root.Left)
		}
		sum += sumOfLeftLeaves(root.Right)
	}
	return sum
}

func isLeaf(root *TreeNode) bool {
	if root == nil || root.Left != nil || root.Right != nil {
		return false
	}
	return true
}
