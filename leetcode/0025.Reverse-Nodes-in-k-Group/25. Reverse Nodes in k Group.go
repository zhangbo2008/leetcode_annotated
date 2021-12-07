package leetcode

import (
	"structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}//node现在是第k+1个,node是下一轮的起点.
	newHead := reverse(head, node)  //反转一个链表,输入的是链表的头和尾. 头到尾的前一个进行反转.
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first *ListNode, last *ListNode) *ListNode {
	prev := last
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}
