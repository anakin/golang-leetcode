package main

var ret int

//直径等于left最大深度+right最大深度+1
func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	right := helper(root.Right)
	if ret < left+right+1 {
		ret = left + right + 1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	ret = 0
	helper(root)
	if ret == 0 {
		return ret
	}
	return ret - 1
}
