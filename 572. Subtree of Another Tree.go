package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return helper(s.Left, t.Left) && helper(s.Right, t.Right)
}
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if helper(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}
