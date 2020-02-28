/*
In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.

Two nodes of a binary tree are cousins if they have the same depth, but have different parents.

We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.

Return true if and only if the nodes corresponding to the values x and y are cousins.


*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root.Val == x || root.Val == y {
		return false
	}
	return bfs([]*TreeNode{root}, x, y)

}

func bfs(nodes []*TreeNode, x, y int) bool {
	if len(nodes) == 0 {
		return false
	}
	xp, yp, next := -1, -1, []*TreeNode{}
	for _, node := range nodes {
		if node.Left != nil {
			if node.Left.Val == x {
				xp = node.Val
			} else if node.Left.Val == y {
				yp = node.Val
			}
			next = append(next, node.Left)
		}
		if node.Right != nil {
			if node.Right.Val == x {
				xp = node.Val
			} else if node.Right.Val == y {
				yp = node.Val
			}
			next = append(next, node.Right)
		}
	}
	if xp != -1 || yp != -1 {
		return xp != yp && xp != -1 && yp != -1
	}
	return bfs(next, x, y)
}

/**
广度优先搜索，xp和yp记录node的parent，判断条件时xp!=yp
*/
