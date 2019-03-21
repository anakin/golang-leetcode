package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
注意题目给的条件
*/
func dfs(root *TreeNode, min int) int {
	if root == nil {
		return -1
	}
	if root.Val > min {
		return root.Val
	}
	left := dfs(root.Left, min)
	right := dfs(root.Right, min)
	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}
	if left < right {
		return left
	}
	return right

}
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return dfs(root, root.Val)
}
