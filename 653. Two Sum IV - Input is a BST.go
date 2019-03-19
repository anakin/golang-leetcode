package main

/**
遍历树，用hash保存节点值，递归判断
类似普通的2sum
*/
var m map[int]bool

func helper(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	if _, ok := m[k-root.Val]; ok {
		return true
	}
	m[root.Val] = true
	return helper(root.Left, k) || helper(root.Right, k)
}
func findTarget(root *TreeNode, k int) bool {
	m = make(map[int]bool)
	return helper(root, k)
}

/**
另外，也可以利用bst的性质，中序遍历成一个有序数组，然后再做2sum
*/
