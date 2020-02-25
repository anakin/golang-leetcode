/*
Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).

The binary search tree is guaranteed to have unique values.



Example 1:

Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
Output: 32
Example 2:

Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
Output: 23

*/
package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func rangeSumBST(root *TreeNode, L int, R int) int {
	res:=0
	help(root,L,R,&res)
	return res
}

func helper(root *TreeNode,L,R int,res *int){
	if root == nil{
		return
	}

	if root.Val>=L && root.Val<=R{
		*res +=root.Val
		helper(root.Left,L,R,res)
		helper(root.Right,L,R,res)
		return
	}
	if root.Val<L{
		helper(root.Right,L,R,res)
	}
	if root.Val>R{
		helper(root.Left,L,R,res)
	}
}
