/**
Given the root of a binary tree, consider all root to leaf paths: paths from the root to any leaf.  (A leaf is a node with no children.)

A node is insufficient if every such root to leaf path intersecting this node has sum strictly less than limit.

Delete all insufficient nodes simultaneously, and return the root of the resulting binary tree.

*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	sufficientSubsetRec(&root, 0, limit)
	return root
}

func sufficientSubsetRec(ptrToCurr **TreeNode, acc int, limit int) {
	cur := *ptrToCurr
	valueSoFar := acc + cur.Val

	hasLeftKid, hasRightKid := cur.Left != nil, cur.Right != nil
	if cur.Left != nil {
		sufficientSubsetRec(&cur.Left, valueSoFar, limit)
	}
	if cur.Right != nil {
		sufficientSubsetRec(&cur.Right, valueSoFar, limit)
	}
	if valueSoFar < limit && cur.Left == nil && cur.Right == nil {
		*ptrToCurr = nil
	}
	if cur.Left == nil && cur.Right == nil && (hasLeftKid || hasRightKid) {
		*ptrToCurr = nil
	}
}
