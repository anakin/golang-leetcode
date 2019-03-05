package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//两个指针，先计算长度，相遇的时候，两个链表的长度一致
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	la, lb := 0, 0
	for a != nil {
		a = a.Next
		la++
	}
	for b != nil {
		b = b.Next
		lb++
	}
	a, b = headA, headB
	for la < lb {
		b = b.Next
		b--
	}
	for lb < la {
		a = a.Next
		a--
	}
	for a != b {
		a = a.Next
		b = b.Next
	}
	return a
}
