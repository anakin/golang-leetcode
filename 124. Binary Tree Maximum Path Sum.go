package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func maxPathSum(root *TreeNode) int {
	ans = math.MinInt32
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	maxPath(root)
	return ans
}
func mymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxPath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := mymax(0, maxPath(root.Left))
	r := mymax(0, maxPath(root.Right))
	sum := l + r + root.Val
	ans = mymax(ans, sum)
	return mymax(l, r) + root.Val

}
