package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	len := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		len++
	}
	tail.Next = head //变成一个环
	k = k % len      //因为是个环，所以不用转超过一圈
	for i := 0; i < len-k; i++ {
		tail = tail.Next
	}
	head = tail.Next
	tail.Next = nil
	return head
}
