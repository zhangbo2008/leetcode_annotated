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

// 解法一       //链表题目用快慢指针.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head} //建立一个val是0的空listnode然后让他指向我们链表的头节点.也就是next赋值为head即可.注意这里面操作都是指针所以要取地址. 还是因为结构体都很大, 复制操作太慢.所以结构体都传指针.
	preSlow, slow, fast := dummyHead, head, head
	for fast != nil {
		if n <= 0 {// 只有n<=0了, slow才开始走. 并且preslow存储每一个slow的前节点.
			preSlow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next  //保证了fast先走n步.
	}//所以当fast到最后的nil时候. slow的前节点就是我们要的东西.
	preSlow.Next = slow.Next// 掰指针即可.删除slow节点.
	return dummyHead.Next//返回链表最开始节点.
}

// 解法二
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n <= 0 {
		return head
	}
	current := head
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	if n > len {
		return head
	}
	if n == len {
		current := head
		head = head.Next
		current.Next = nil
		return head
	}
	current = head
	i := 0
	for current != nil {
		if i == len-n-1 {
			deleteNode := current.Next
			current.Next = current.Next.Next
			deleteNode.Next = nil
			break
		}
		i++
		current = current.Next
	}
	return head
}
