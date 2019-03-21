package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret = TreeNode{Val: 0}
var tmp = &ret

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	increasingBST(root.Left)
	tmp.Right = &TreeNode{
		Val: root.Val,
	}
	tmp = tmp.Right
	increasingBST(root.Right)
	return ret.Right
}
