package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newh ListNode
	for head != nil {
		newh.Next, head.Next, head = head, newh.Next, head.Next
	}
	return newh.Next

}
