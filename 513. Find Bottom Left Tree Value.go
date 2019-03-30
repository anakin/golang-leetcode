package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
深度优先搜索整个树，记录层数
每次搜索的时候，第一个碰到的元素，就是该层的最左子树
*/
var maxd, val int

func findBottomLeftValue(root *TreeNode) int {
	maxd = 0
	val = 0
	dfs(root, 1)
	return val
}

func dfs(root *TreeNode, d int) {
	if root == nil {
		return
	}
	if d > maxd {
		maxd = d
		val = root.Val
	}
	if root.Left != nil {
		dfs(root.Left, d+1)
	}
	if root.Right != nil {
		dfs(root.Right, d+1)
	}

}
