package main

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
