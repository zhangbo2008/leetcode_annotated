# [1642. Furthest Building You Can Reach](https://leetcode.com/problems/furthest-building-you-can-reach/)


## 题目

You are given an integer array `heights` representing the heights of buildings, some `bricks`, and some `ladders`.

You start your journey from building `0` and move to the next building by possibly using bricks or ladders.

While moving from building `i` to building `i+1` (**0-indexed**),

- If the current building's height is **greater than or equal** to the next building's height, you do **not** need a ladder or bricks.
- If the current building's height is **less than** the next building's height, you can either use **one ladder** or `(h[i+1] - h[i])` **bricks**.

*Return the furthest building index (0-indexed) you can reach if you use the given ladders and bricks optimally.*

**Example 1:**

![https://assets.leetcode.com/uploads/2020/10/27/q4.gif](https://assets.leetcode.com/uploads/2020/10/27/q4.gif)

```
Input: heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1
Output: 4
Explanation: Starting at building 0, you can follow these steps:
- Go to building 1 without using ladders nor bricks since 4 >= 2.
- Go to building 2 using 5 bricks. You must use either bricks or ladders because 2 < 7.
- Go to building 3 without using ladders nor bricks since 7 >= 6.
- Go to building 4 using your only ladder. You must use either bricks or ladders because 6 < 9.
It is impossible to go beyond building 4 because you do not have any more bricks or ladders.

```

**Example 2:**

```
Input: heights = [4,12,2,7,3,18,20,3,19], bricks = 10, ladders = 2
Output: 7

```

**Example 3:**

```
Input: heights = [14,3,19,3], bricks = 17, ladders = 0
Output: 3

```

**Constraints:**

- `1 <= heights.length <= 10^5`
- `1 <= heights[i] <= 10^6`
- `0 <= bricks <= 10^9`
- `0 <= ladders <= heights.length`

## 题目大意

给你一个整数数组 heights ，表示建筑物的高度。另有一些砖块 bricks 和梯子 ladders 。你从建筑物 0 开始旅程，不断向后面的建筑物移动，期间可能会用到砖块或梯子。当从建筑物 i 移动到建筑物 i+1（下标 从 0 开始 ）时：

- 如果当前建筑物的高度 大于或等于 下一建筑物的高度，则不需要梯子或砖块。
- 如果当前建筑的高度 小于 下一个建筑的高度，您可以使用 一架梯子 或 (h[i+1] - h[i]) 个砖块

如果以最佳方式使用给定的梯子和砖块，返回你可以到达的最远建筑物的下标（下标 从 0 开始 ）。

## 解题思路

- 这一题可能会想到贪心算法。梯子很厉害，可以无限长，所以梯子用来跨越最高的楼。遇到非最高的距离差，先用砖头。这样贪心的话不正确。例如，[1, 5, 1, 2, 3, 4, 10000] 这组数据，梯子有 1 个，4 块砖头。最大的差距在 10000 和 4 之间，贪心选择在此处用梯子。但是砖头不足以让我们走到最后两栋楼。贪心得到的结果是 3，正确的结果是 5，先用梯子，再用砖头走过 3，4，5 号楼。
- 上面的贪心解法错误在于没有“动态”的贪心，使用梯子应该选择能爬过楼里面最高的 2 个。于是顺理成章的想到了优先队列。维护一个长度为梯子个数的最小堆，当队列中元素超过梯子个数，便将队首最小值出队，出队的这个楼与楼的差距用砖头填补。所有砖头用完了，即是可以到达的最远楼号。

## 代码

```go
package leetcode

import (
	"container/heap"
)

func furthestBuilding(heights []int, bricks int, ladder int) int {
	usedLadder := &heightDiffPQ{}
	for i := 1; i < len(heights); i++ {
		needbricks := heights[i] - heights[i-1]
		if needbricks < 0 {
			continue
		}
		if ladder > 0 {
			heap.Push(usedLadder, needbricks)
			ladder--
		} else {
			if len(*usedLadder) > 0 && needbricks > (*usedLadder)[0] {
				needbricks, (*usedLadder)[0] = (*usedLadder)[0], needbricks
				heap.Fix(usedLadder, 0)
			}
			if bricks -= needbricks; bricks < 0 {
				return i - 1
			}
		}
	}
	return len(heights) - 1
}

type heightDiffPQ []int

func (pq heightDiffPQ) Len() int            { return len(pq) }
func (pq heightDiffPQ) Less(i, j int) bool  { return pq[i] < pq[j] }
func (pq heightDiffPQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *heightDiffPQ) Push(x interface{}) { *pq = append(*pq, x.(int)) }
func (pq *heightDiffPQ) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return x
}
```