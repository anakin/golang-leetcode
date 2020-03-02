/*
We are given a linked list with head as the first node.  Let's number the nodes in the list: node_1, node_2, node_3, ... etc.

Each node may have a next larger value: for node_i, next_larger(node_i) is the node_j.val such that j > i, node_j.val > node_i.val, and j is the smallest possible choice.  If such a j does not exist, the next larger value is 0.

Return an array of integers answer, where answer[i] = next_larger(node_{i+1}).

Note that in the example inputs (not outputs) below, arrays such as [2,1,5] represent the serialization of a linked list with a head node value of 2, second node value of 1, and third node value of 5.



Example 1:

Input: [2,1,5]
Output: [5,5,0]
Example 2:

Input: [2,7,4,3,5]
Output: [7,0,5,5,0]
Example 3:

Input: [1,7,5,1,9,2,5,1]
Output: [7,9,9,9,0,5,0,0]

*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	var stack []int
	top := -1
	stack = append(stack, 0)
	top++
	for i := 1; i < len(nums); i++ {
		for top != -1 && nums[stack[top]] < nums[i] {
			nums[stack[top]] = nums[i]
			stack = stack[:top]
			top--
		}
		stack = append(stack, i)
		top++
	}
	for top != -1 {
		nums[stack[top]] = 0
		stack = stack[:top]
		top--
	}
	return nums
}

/**
维护一个单调递减的栈，当前元素小于栈顶元素时，就是要找的元素，出栈
*/
