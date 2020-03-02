/**
Given the root of a binary search tree with distinct values, modify it so that every node has a new value equal to the sum of the values of the original tree that are greater than or equal to node.val.

As a reminder, a binary search tree is a tree that satisfies these constraints:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	helper(root, nil)
	return root
}
func helper(root, max *TreeNode) {
	if root == nil {
		return
	}
	root.Val += findSum(root.Right)
	if max != nil {
		root.Val += max.Val
	}
	helper(root.Left, root)
	helper(root.Right, max)
}

func findSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := root.Val
	sum += findSum(root.Left)
	sum += findSum(root.Right)
	return sum
}
