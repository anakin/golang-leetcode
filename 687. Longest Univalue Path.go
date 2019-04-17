package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func longestUnivaluePath(root *TreeNode) int {
	ans = 0
	longestPath(root)
	return ans
}

func longestPath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := longestPath(root.Left)
	r := longestPath(root.Right)
	pl, pr := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		pl = l + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		pr = r + 1
	}
	ans = mymax(ans, pl+pr)
	return mymax(pl, pr)
}

func mymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
