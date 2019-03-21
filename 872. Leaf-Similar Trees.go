package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
先递归找到所有的叶子节点，再循环判断
*/

func leaf(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*ret = append(*ret, root.Val)
	}
	leaf(root.Left, ret)
	leaf(root.Right, ret)
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	a, b := []int{}, []int{}
	leaf(root1, &a)
	leaf(root2, &b)
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
