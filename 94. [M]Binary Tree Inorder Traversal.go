package main

/**
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	helper(&res, root)
	return res
}
func helper(res *[]int, root *TreeNode) {
	if root.Left != nil {
		helper(res, root.Left)
	}
	*res = append(*res, root.Val)
	if root.Right != nil {
		helper(res, root.Right)
	}

}
