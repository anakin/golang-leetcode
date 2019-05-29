package main

/**
Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

Example 1:

Input: 1->2->3->3->4->4->5
Output: 1->2->5
Example 2:

Input: 1->1->1->2->3
Output: 2->3
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	t := &ListNode{}
	t.Next = head
	p := t
	for p.Next != nil && p.Next.Next != nil {
		if p.Next.Val == p.Next.Next.Val {
			dup := p.Next.Val
			for p.Next != nil && p.Next.Val == dup {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}
	return t.Next
}
