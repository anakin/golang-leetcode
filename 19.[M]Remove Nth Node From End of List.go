package main

/**
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.
*/

/**
* 快慢指针,快指针先前进n步，然后两个指针一起前进，当快指针走到末尾的时候，慢指针所在的位置，就是目标位置
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

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
