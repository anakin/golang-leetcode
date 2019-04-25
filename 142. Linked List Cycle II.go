package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
用快慢指针的方式，快指针每次走两步，慢指针每次走一步，相遇的时候，就是环的位置
此时slow停的位置到环入口的距离，等于头结点到环入口的距离（https://blog.csdn.net/jhlovetll/article/details/85255708）
所以两个结点前进相同的步数，head就位于环的入口处了
*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head.Next, head.Next.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow, fast = slow.Next, fast.Next.Next
	}
	//没有环
	if slow != fast {
		return nil
	}
	for slow != head {
		slow, head = slow.Next, head.Next
	}
	return slow
}
