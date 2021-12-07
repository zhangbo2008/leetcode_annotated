# [987. Vertical Order Traversal of a Binary Tree](https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/)


## 题目

Given the `root` of a binary tree, calculate the **vertical order traversal** of the binary tree.

For each node at position `(row, col)`, its left and right children will be at positions `(row + 1, col - 1)` and `(row + 1, col + 1)` respectively. The root of the tree is at `(0, 0)`.

The **vertical order traversal** of a binary tree is a list of top-to-bottom orderings for each column index starting from the leftmost column and ending on the rightmost column. There may be multiple nodes in the same row and same column. In such a case, sort these nodes by their values.

Return *the **vertical order traversal** of the binary tree*.

**Example 1:**

![https://assets.leetcode.com/uploads/2021/01/29/vtree1.jpg](https://assets.leetcode.com/uploads/2021/01/29/vtree1.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: [[9],[3,15],[20],[7]]
Explanation:
Column -1: Only node 9 is in this column.
Column 0: Nodes 3 and 15 are in this column in that order from top to bottom.
Column 1: Only node 20 is in this column.
Column 2: Only node 7 is in this column.
```

**Example 2:**

![https://assets.leetcode.com/uploads/2021/01/29/vtree2.jpg](https://assets.leetcode.com/uploads/2021/01/29/vtree2.jpg)

```
Input: root = [1,2,3,4,5,6,7]
Output: [[4],[2],[1,5,6],[3],[7]]
Explanation:
Column -2: Only node 4 is in this column.
Column -1: Only node 2 is in this column.
Column 0: Nodes 1, 5, and 6 are in this column.
          1 is at the top, so it comes first.
          5 and 6 are at the same position (2, 0), so we order them by their value, 5 before 6.
Column 1: Only node 3 is in this column.
Column 2: Only node 7 is in this column.

```

**Example 3:**

![https://assets.leetcode.com/uploads/2021/01/29/vtree3.jpg](https://assets.leetcode.com/uploads/2021/01/29/vtree3.jpg)

```
Input: root = [1,2,3,4,6,5,7]
Output: [[4],[2],[1,5,6],[3],[7]]
Explanation:
This case is the exact same as example 2, but with nodes 5 and 6 swapped.
Note that the solution remains the same since 5 and 6 are in the same location and should be ordered by their values.

```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 1000]`.
- `0 <= Node.val <= 1000`

## 题目大意

给你二叉树的根结点 root ，请你设计算法计算二叉树的 垂序遍历 序列。

对位于 (row, col) 的每个结点而言，其左右子结点分别位于 (row + 1, col - 1) 和 (row + 1, col + 1) 。树的根结点位于 (0, 0) 。二叉树的 垂序遍历 从最左边的列开始直到最右边的列结束，按列索引每一列上的所有结点，形成一个按出现位置从上到下排序的有序列表。如果同行同列上有多个结点，则按结点的值从小到大进行排序。返回二叉树的 垂序遍历 序列。

## 解题思路

- 题目要求按照一列一列的遍历二叉树。需要解决 2 个问题。第一个问题，二叉树上每个结点的二维坐标如何计算。第二个问题，同一个二维坐标点上摞起来多个结点，需要按照从小到大的顺序排序，如例子二和例子三，同一个二维坐标点 (2，0) 上，摞了 2 个不同的结点。
- 先解决第一个问题，由于题目要求根结点是 (0，0) ，即根结点是坐标原点，它的左子树的 x 坐标都是负数，它的右子树的 x 坐标都是正数。按照先序遍历，就可以将这些结点的二维坐标计算出来。再进行一次排序，按照 x 坐标从小到大排序，坐标相同的情况对应着结点摞起来的情况，摞起来的结点按照 val 值的大小从小到大排序。这样在 x 轴方向，所有结点就排列好了。排序完成，也顺便解决了第二个问题。
- 最后一步只需要扫描一遍这个排好序的数组，按照列的顺序，依次将同一列的结点打包至一个一维数组中。最终输出的二维数组即为题目所求。

## 代码

```go
package leetcode

import (
	"sort"

	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type node struct {
	x, y, val int
}

func verticalTraversal(root *TreeNode) [][]int {
	nodes := []*node{}
	inorder(root, 0, 0, &nodes)
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].y == nodes[j].y {
			if nodes[i].x < nodes[j].x {
				return true
			} else if nodes[i].x > nodes[j].x {
				return false
			}
			return nodes[i].val < nodes[j].val
		}
		return nodes[i].y < nodes[j].y
	})
	res, currY, currColumn := [][]int{}, nodes[0].y, []int{nodes[0].val}
	for i := 1; i < len(nodes); i++ {
		if currY == nodes[i].y {
			currColumn = append(currColumn, nodes[i].val)
		} else {
			res = append(res, currColumn)
			currColumn = []int{nodes[i].val}
			currY = nodes[i].y
		}
	}
	res = append(res, currColumn)
	return res
}

func inorder(root *TreeNode, x, y int, nodes *[]*node) {
	if root != nil {
		*nodes = append(*nodes, &node{x, y, root.Val})
		inorder(root.Left, x+1, y-1, nodes)
		inorder(root.Right, x+1, y+1, nodes)
	}
}
```