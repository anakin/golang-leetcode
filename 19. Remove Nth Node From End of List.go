package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 快慢指针,快指针先前进n步，然后两个指针一起前进，当快指针走到末尾的时候，慢指针所在的位置，就是目标位置
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
