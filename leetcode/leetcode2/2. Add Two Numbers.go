package leetcode2

import (
	"awesomeProject1/structures"
)

// ListNode define
type ListNode = structures.ListNode   //listnode的结构,

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head  //carry 表示是否发生进位
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10} //当前节点是余数
		current = current.Next //然后走下一个节点
		carry = (n1 + n2 + carry) / 10 //carry是更新近卫数.
	}
	return head.Next
}
