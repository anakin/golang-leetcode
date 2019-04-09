package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
递归实现
*/

var sum int

func sumNumbers(root *TreeNode) int {
	sum = 0
	if root == nil {
		return 0
	}
	helper(root, 0)
	return sum
}

func helper(root *TreeNode, num int) int {
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		sum += num
	}
	if root.Left != nil {
		helper(root.Left, num)
	}
	if root.Right != nil {
		helper(root.Right, num)
	}
}
