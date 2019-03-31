package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	shadowHead := &ListNode{Next: head} //用于返回
	prevNode := shadowHead
	//走到反转的起点位置
	for i := 1; i < m; i++ {
		prevNode = prevNode.Next
		head = head.Next
	}
	mNode := head
	nNode := head.Next
	//开始反转操作
	for i := m; i < n; i++ {
		nextNNode := nNode.Next
		nNode.Next = head
		head = nNode
		nNode = nextNNode
	}
	//把反转之后的list和原来的对接
	mNode.Next = nNode
	prevNode.Next = head
	return shadowHead.Next
}
