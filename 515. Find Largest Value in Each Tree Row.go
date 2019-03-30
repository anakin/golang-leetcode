package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
前序遍历，先把每层的第一个节点加入结果数组，然后判断每个节点和第一个节点的大小
*/
var res []int

func largestValues(root *TreeNode) []int {
	res = []int{}
	preorder(root, 0)
	return res
}

func preorder(root *TreeNode, d int) {
	if root == nil {
		return
	}
	if d == len(res) {
		res = append(res, root.Val)
	} else {
		res[d] = max(res[d], root.Val)
	}
	preorder(root.Left, d+1)
	preorder(root.Right, d+1)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
