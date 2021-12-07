# [1669. Merge In Between Linked Lists](https://leetcode.com/problems/merge-in-between-linked-lists/)


## 题目

You are given two linked lists: `list1` and `list2` of sizes `n` and `m` respectively.

Remove `list1`'s nodes from the `ath` node to the `bth` node, and put `list2` in their place.

The blue edges and nodes in the following figure incidate the result:

![https://assets.leetcode.com/uploads/2020/11/05/fig1.png](https://assets.leetcode.com/uploads/2020/11/05/fig1.png)

*Build the result list and return its head.*

**Example 1:**

![https://assets.leetcode.com/uploads/2020/11/05/merge_linked_list_ex1.png](https://assets.leetcode.com/uploads/2020/11/05/merge_linked_list_ex1.png)

```
Input: list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
Output: [0,1,2,1000000,1000001,1000002,5]
Explanation: We remove the nodes 3 and 4 and put the entire list2 in their place. The blue edges and nodes in the above figure indicate the result.

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/11/05/merge_linked_list_ex2.png](https://assets.leetcode.com/uploads/2020/11/05/merge_linked_list_ex2.png)

```
Input: list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
Output: [0,1,1000000,1000001,1000002,1000003,1000004,6]
Explanation: The blue edges and nodes in the above figure indicate the result.

```

**Constraints:**

- `3 <= list1.length <= 104`
- `1 <= a <= b < list1.length - 1`
- `1 <= list2.length <= 104`

## 题目大意

给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。请你将 list1 中第 a 个节点到第 b 个节点删除，并将list2 接在被删除节点的位置。

## 解题思路

- 简单题，考查链表的基本操作。此题注意 a == b 的情况。

## 代码

```go
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	n := list1
	var startRef, endRef *ListNode
	for i := 0; i <= b; i++ {
		if i == a-1 {
			startRef = n
		}
		if i == b {
			endRef = n
		}
		n = n.Next
	}
	startRef.Next = list2
	n = list2
	for n.Next != nil {
		n = n.Next
	}
	n.Next = endRef.Next
	return list1
}
```