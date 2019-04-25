package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
用一个slice实现stack结构，首先将root和root的左子树入栈
出栈的同时，将右子树入栈
*/

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := make([]*TreeNode, 0, 128)
	res := BSTIterator{
		stack: stack,
	}
	res.push(root)
	return res
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	size := len(this.stack)
	var top *TreeNode
	this.stack, top = this.stack[:size-1], this.stack[size-1]
	this.push(top.Right)
	return top.Val

}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

func (this *BSTIterator) push(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}
