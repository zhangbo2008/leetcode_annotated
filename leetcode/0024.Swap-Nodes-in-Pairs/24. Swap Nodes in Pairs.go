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

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	//根据题意,每次需要算4个节点.  然后算好这4个节点的next指针修改好之后, 进行头指针移动.
	// 这里面的头指针移动是 pt=pt.Next这个操作.
	for pt := dummy; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return dummy.Next
}
