package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
把链表的前半部分反转，如果整个链表是回文的话，反转之后应该是
1->2->1->2这样的形式
然后依次比较对应位置上的值是否相等
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow, tmp := head, head, head
	var l ListNode
	for fast != nil && fast.Next != nil {
		//快指针每次前进两步
		fast = fast.Next.Next
		//慢指针负责反转前半部分链表
		tmp = slow.Next
		slow.Next = l.Next
		l.Next = slow
		slow = tmp
	}
	if fast != nil {
		slow = slow.Next
	}
	tmp = l.Next
	//比较对应位置上的节点值是否相等
	for slow != nil && tmp != nil && slow.Val == tmp.Val {
		slow, tmp = slow.Next, tmp.Next
	}
	return slow == nil
}
