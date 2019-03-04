package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//快慢指针，如果相遇则有环
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
