package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	s := []string{}
	if root == nil {
		return s
	}
	helper(root, strconv.Itoa(root.Val), &s)
	return s
}

func helper(node *TreeNode, path string, s *[]string) {
	if node.Left == nil && node.Right == nil {
		*s = append(*s, path)
	}
	if node.Left != nil {
		helper(node.Left, path+"->"+strconv.Itoa(node.Left.Val), s)
	}
	if node.Right != nil {
		helper(node.Right, path+"->"+strconv.Itoa(node.Right.Val), s)
	}

}
