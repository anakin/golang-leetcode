package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
先计算左子树中节点的个数
如果左子树节点个数大于k-1，那么答案在左子树中
如果个数小于k-1，就在右子树中
如果刚好相等，就直接返回root节点的值
*/

func kthSmallest(root *TreeNode, k int) int {
	left := countLeaf(root.Left)
	if left == k-1 {
		return root.Val
	} else if left > k-1 {
		return kthSmallest(root.Left, k)
	} else {
		return kthSmallest(root.Right, k-left-1)
	}
}

func countLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countLeaf(root.Left) + countLeaf(root.Right) + 1
}
