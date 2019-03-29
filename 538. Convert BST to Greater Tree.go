package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
因为给定的树是BST，所以可以知道如果采用后续遍历的方式，第一个找到的节点应该是最大的节点。所以采用递归的方式从最大的节点开始改变树的值
*/
var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	helper(root)
	return root
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Right)

	root.Val += sum
	sum = root.Val
	helper(root.Left)
}
