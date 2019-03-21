package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
先中序遍历，得到一个有序的数组
*/

func helper(root *TreeNode, l *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, l)
	*l = append(*l, root.Val)
	helper(root.Right, l)
}

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var l []int
	helper(root, &l)
	ret := math.MaxInt32
	for i := 1; i < len(l); i++ {
		if l[i]-l[i-1] < ret {
			ret = l[i] - l[i-1]
		}
	}
	return ret
}
