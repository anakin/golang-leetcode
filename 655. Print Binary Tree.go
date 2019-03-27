package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
先得到树的最大高度h
打印的宽度等于2的h次方-1
*/

var ret [][]string

func printTree(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}
	h := getHeight(root)
	w := pow(2, h) - 1
	ret = make([][]string, h)
	for k := range ret {
		s := make([]string, w)
		for key := range s {
			s[key] = ""
		}
		ret[k] = s
	}
	helper(root, 0, 0, w)
	return ret
}
func helper(root *TreeNode, level, l, r int) {
	if root != nil {
		mid := l + (r-l)/2
		ret[level][mid] = strconv.Itoa(root.Val)
		helper(root.Left, level+1, l, mid-1)
		helper(root.Right, level+1, mid+1, r)
	}
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := getHeight(root.Left)
	r := getHeight(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
func pow(x, y int) int {
	ret := x
	for i := 1; i < y; i++ {
		ret *= x
	}
	return ret
}
