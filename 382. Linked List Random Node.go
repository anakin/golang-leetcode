package main

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	count int
	head  *ListNode
	pool  []int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	var s Solution
	s.head = head
	return s
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	r := rand.Intn(this.count + 1)
	this.count++
	if r == this.count-1 && this.head != nil {
		t := this.head
		this.head = this.head.Next
		this.pool = append(this.pool, t.Val)
		return t.Val
	} else {
		return this.pool[r%len(this.pool)]
	}

}
