package main

//双指针
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	var p1 ListNode
	p1.Next = head
	p2 := &p1
	for head != nil {
		if head.Val == val {
			p2.Next, head = head.Next, head.Next
		} else {
			p2, head = head, head.Next
		}
	}
	return p1.Next
}
