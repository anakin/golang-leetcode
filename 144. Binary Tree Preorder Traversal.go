package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前中后基本类似
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	helper(&res, root)
	return res
}

func helper(res *[]int, root *TreeNode) {
	*res = append(*res, root.Val)
	if root.Left != nil {
		helper(res, root.Left)
	}
	if root.Right != nil {
		helper(res, root.Right)
	}
}
