package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val { //如果key小于root，从左子树删除
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val { //如果key大于root，从右子树删除
		root.Right = deleteNode(root.Right, key)
	} else { //删除root节点
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		root.Val = findMin(root.Right)
		root.Right = deleteNode(root.Right, root.Val)
	}
	return root
}
func findMin(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}
