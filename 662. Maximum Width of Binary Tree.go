package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int
var left map[int]int

func widthOfBinaryTree(root *TreeNode) int {
	ans = 0
	left = make(map[int]int)
	dfs(root, 0, 0)
	return ans
}
func dfs(root *TreeNode, depth, pos int) {
	if root == nil {
		return
	}
	if _, ok := left[depth]; !ok {
		left[depth] = pos
	}
	ans = max(ans, pos-left[depth]+1)
	dfs(root.Left, depth+1, 2*pos)
	dfs(root.Right, depth+1, 2*pos+1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}
