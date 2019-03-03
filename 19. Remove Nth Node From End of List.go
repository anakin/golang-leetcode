package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 快慢指针
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	last := head
	for i := 0; i < n; i++ {
		if cur.Next != nil {
			cur = cur.Next
		} else {
			return head.Next
		}
	}
	for cur.Next != nil {
		cur = cur.Next
		last = last.Next
	}
	last.Next = last.Next.Next
	return head
}
