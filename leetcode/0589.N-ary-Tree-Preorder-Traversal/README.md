# [589. N-ary Tree Preorder Traversal](https://leetcode.com/problems/n-ary-tree-preorder-traversal/)

## 题目

Given the `root` of an n-ary tree, return *the preorder traversal of its nodes' values*.

Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value (See examples)

**Example 1:**

![https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png](https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png)

```
Input: root = [1,null,3,2,4,null,5,6]
Output: [1,3,5,6,2,4]
```

**Example 2:**

![https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png](https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png)

```
Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 104]`.
- `0 <= Node.val <= 10^4`
- The height of the n-ary tree is less than or equal to `1000`.

**Follow up:** Recursive solution is trivial, could you do it iteratively?

## 题目大意

给定一个 N 叉树，返回其节点值的 **前序遍历** 。N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 `null` 分隔（请参见示例）。

## 解题思路

- N 叉树和二叉树的前序遍历原理完全一样。二叉树非递归解法需要用到栈辅助，N 叉树同样如此。将父节点的所有孩子节点**逆序**入栈，逆序的目的是为了让前序节点永远在栈顶。依次循环直到栈里所有元素都出栈。输出的结果即为 N 叉树的前序遍历。时间复杂度 O(n)，空间复杂度 O(n)。
- 递归解法非常简单，见解法二。

## 代码

```go
package leetcode

//  Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// 解法一 非递归
func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, r.Val)
		tmp := []*Node{}
		for _, v := range r.Children {
			tmp = append([]*Node{v}, tmp...) // 逆序存点
		}
		stack = append(stack, tmp...)
	}
	return res
}

// 解法二 递归
func preorder1(root *Node) []int {
	res := []int{}
	preorderdfs(root, &res)
	return res
}

func preorderdfs(root *Node, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		for i := 0; i < len(root.Children); i++ {
			preorderdfs(root.Children[i], res)
		}
	}
}
```