package main

/**
节点顺序入栈
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1, len2 := 0, 0
	for tmp := l1; tmp != nil; tmp = tmp.Next {
		len1++
	}
	for tmp := l2; tmp != nil; tmp = tmp.Next {
		len2++
	}
	if len1 < len2 {
		l1, l2, len1, len2 = l2, l1, len2, len1
	}
	s1, s2 := make([]*ListNode, 0, len1), make([]*ListNode, 0, len2)
	for tmp := l1; tmp != nil; tmp = tmp.Next {
		s1 = append(s1, tmp)
	}
	for tmp := l2; tmp != nil; tmp = tmp.Next {
		s2 = append(s2, tmp)
	}

	var head ListNode
	add, i, t := 0, 0, 0
	for i = 1; i <= len1; i++ {
		if i <= len2 {
			t = s1[len1-i].Val + s2[len2-i].Val + add
		} else {
			t = s1[len1-i].Val + add
		}
		add = t / 10
		node := &ListNode{
			t % 10,
			head.Next,
		}
		head.Next = node
	}
	if add > 0 {
		node := &ListNode{
			add,
			head.Next,
		}
		head.Next = node
	}
	return head.Next
}
