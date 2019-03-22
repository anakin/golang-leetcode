package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
递归判断
*/
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Val != root.Left.Val {
		return false
	}
	if root.Right != nil && root.Val != root.Right.Val {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
