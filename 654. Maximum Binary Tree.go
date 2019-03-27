package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
递归
先找到最大的，作为root
然后左边的作为Left，右边的作为Right

*/

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return helper(nums, 0, len(nums))
}

func helper(nums []int, l, r int) *TreeNode {
	if l == r {
		return nil
	}
	maxi := max(nums, l, r)
	root := &TreeNode{Val: nums[maxi]}
	root.Left = helper(nums, l, maxi)
	root.Right = helper(nums, maxi+1, r)
	return root
}

func max(nums []int, l, r int) int {
	ret := l
	for i := l; i < r; i++ {
		if nums[ret] < nums[i] {
			ret = i
		}
	}
	return ret
}
