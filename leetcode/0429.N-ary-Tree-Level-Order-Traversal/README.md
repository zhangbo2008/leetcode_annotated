# [429. N-ary Tree Level Order Traversal](https://leetcode.com/problems/n-ary-tree-level-order-traversal/)


## 题目

Given an n-ary tree, return the *level order* traversal of its nodes' values.

*Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).*

**Example 1:**

![https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png](https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png)

```
Input: root = [1,null,3,2,4,null,5,6]
Output: [[1],[3,2,4],[5,6]]

```

**Example 2:**

![https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png](https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png)

```
Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]

```

**Constraints:**

- The height of the n-ary tree is less than or equal to `1000`
- The total number of nodes is between `[0, 104]`

## 题目大意

给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

## 解题思路

- 这是 n 叉树的系列题，第 589 题也是这一系列的题目。这一题思路不难，既然是层序遍历，用 BFS 解答。

## 代码

```go
package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var res [][]int
	var temp []int
	if root == nil {
		return res
	}
	queue := []*Node{root, nil}
	for len(queue) > 1 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			queue = append(queue, nil)
			res = append(res, temp)
			temp = []int{}
		} else {
			temp = append(temp, node.Val)
			if len(node.Children) > 0 {
				queue = append(queue, node.Children...)
			}
		}
	}
	res = append(res, temp)
	return res
}
```