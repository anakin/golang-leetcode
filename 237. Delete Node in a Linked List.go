package main

type ListNode struct {
	Val  int
	Next *TreeNode
}

func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	if node.Next == nil {
		node = nil
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
