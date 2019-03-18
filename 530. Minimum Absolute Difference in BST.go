package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var st []int

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorder(root.Left)
	}
	if root != nil {
		st = append(st, root.Val)
	}
	if root.Right != nil {
		inorder(root.Right)
	}
}

func getMinimumDifference(root *TreeNode) int {
	st = make([]int, 0, 256)
	inorder(root)
	ret := math.MaxInt32
	for i := 1; i < len(st); i++ {
		if (st[i] - st[i-1]) < ret {
			ret = st[i] - st[i-1]
		}
	}
	return ret
}
