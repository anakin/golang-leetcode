package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret int

func sumRootToLeaf(root *TreeNode) int {
	ret = 0
	helper(root, 0)
	return ret
}
func helper(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum = sum << 1
	sum += root.Val
	if root.Left == nil && root.Right == nil {
		ret += sum
		return
	}
	if root.Left != nil {
		helper(root.Left, sum)
	}
	if root.Right != nil {
		helper(root.Right, sum)
	}
}
