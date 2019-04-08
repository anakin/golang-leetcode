package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
递归
分两种情况，如果包括root，那么接下来就是root.left.left，root.left.right
如果不包括root，接下来就是root.left,root.right
取两种情况的最大值
*/
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	val := 0
	if root.Left != nil {
		val += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		val += rob(root.Right.Right) + rob(root.Right.Left)
	}
	return max(val+root.Val, rob(root.Left)+rob(root.Right))
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
