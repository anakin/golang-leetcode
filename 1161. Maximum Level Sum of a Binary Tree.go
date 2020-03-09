/**
Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.

Return the smallest level X such that the sum of all the values of nodes at level X is maximal.

Example 1:

Input: [1,7,0,7,-8,null,null]
Output: 2
Explanation:
Level 1 sum = 1.
Level 2 sum = 7 + 0 = 7.
Level 3 sum = 7 + -8 = -1.
So we return the level with the maximum sum which is level 2.

*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m = make(map[int]int)

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dfs(root, 1)
	level := 0
	for l := range m {
		if m[l] > m[level] {
			level = 1
		}
	}
	return level
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	dfs(root.Left, level+1)
	m[level] += root.Val
	dfs(root.Right, level+1)
}
