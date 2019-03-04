package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//新建两个列表
func partition(head *ListNode, x int) *ListNode {

	beforehead := &ListNode{}
	before := beforehead
	afterhead := &ListNode{}
	after := afterhead
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterhead.Next
	return beforehead.Next
}
