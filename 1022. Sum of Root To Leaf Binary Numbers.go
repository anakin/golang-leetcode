/*
Given a binary tree, each node has value 0 or 1.  Each root-to-leaf path represents a binary number starting with the most significant bit.  For example, if the path is 0 -> 1 -> 1 -> 0 -> 1, then this could represent 01101 in binary, which is 13.

For all leaves in the tree, consider the numbers represented by the path from the root to that leaf.

Return the sum of these numbers.

Example 1:
Input: [1,0,1,0,1,0,1]
Output: 22
Explanation: (100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret int

func sumRootToLeaf(root *TreeNode) int {
	ret = 0
	helper(root, 0)
	return ret
}
func helper(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum = sum << 1
	sum += root.Val
	if root.Left == nil && root.Right == nil {
		ret += sum
		return
	}
	if root.Left != nil {
		helper(root.Left, sum)
	}
	if root.Right != nil {
		helper(root.Right, sum)
	}
}
