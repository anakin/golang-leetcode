package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
注意的问题：
1.进位
2.两个list的长度不一定相等
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0, nil}
	cursor := result
	leftBit, rightBit, carryBit := 0, 0, 0

	for l1 != nil || l2 != nil || carryBit > 0 {
		if l1 != nil {
			leftBit = l1.Val
			l1 = l1.Next
		} else {
			leftBit = 0
		}

		if l2 != nil {
			rightBit = l2.Val
			l2 = l2.Next
		} else {
			rightBit = 0
		}
		//考虑有进位的情况
		cursor.Val = (leftBit + rightBit + carryBit) % 10
		carryBit = (leftBit + rightBit + carryBit) / 10

		if l1 != nil || l2 != nil || carryBit > 0 {
			cursor.Next = &ListNode{0, nil}
			cursor = cursor.Next
		}
	}

	return result
}
