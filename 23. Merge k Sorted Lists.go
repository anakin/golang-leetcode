package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
分治的方法，不停的两两merge
其中merge两个list的方法，就是#21
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	begin, end := 0, len(lists)-1
	for begin < end {
		mid := (begin + end - 1) / 2
		for i := 0; i <= mid; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[end-i])
		}
		end = (begin + end) / 2
	}
	return lists[0]
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}

}
