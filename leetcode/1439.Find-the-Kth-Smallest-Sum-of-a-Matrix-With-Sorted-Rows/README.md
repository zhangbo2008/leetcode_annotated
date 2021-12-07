# [1439. Find the Kth Smallest Sum of a Matrix With Sorted Rows](https://leetcode.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/)


## 题目

You are given an `m * n` matrix, `mat`, and an integer `k`, which has its rows sorted in non-decreasing order.

You are allowed to choose exactly 1 element from each row to form an array. Return the Kth **smallest** array sum among all possible arrays.

**Example 1:**

```
Input: mat = [[1,3,11],[2,4,6]], k = 5
Output: 7
Explanation: Choosing one element from each row, the first k smallest sum are:
[1,2], [1,4], [3,2], [3,4], [1,6]. Where the 5th sum is 7.  
```

**Example 2:**

```
Input: mat = [[1,3,11],[2,4,6]], k = 9
Output: 17
```

**Example 3:**

```
Input: mat = [[1,10,10],[1,4,5],[2,3,6]], k = 7
Output: 9
Explanation: Choosing one element from each row, the first k smallest sum are:
[1,1,2], [1,1,3], [1,4,2], [1,4,3], [1,1,6], [1,5,2], [1,5,3]. Where the 7th sum is 9.  
```

**Example 4:**

```
Input: mat = [[1,1,10],[2,2,9]], k = 7
Output: 12
```

**Constraints:**

- `m == mat.length`
- `n == mat.length[i]`
- `1 <= m, n <= 40`
- `1 <= k <= min(200, n ^ m)`
- `1 <= mat[i][j] <= 5000`
- `mat[i]` is a non decreasing array.

## 题目大意

给你一个 m * n 的矩阵 mat，以及一个整数 k ，矩阵中的每一行都以非递减的顺序排列。你可以从每一行中选出 1 个元素形成一个数组。返回所有可能数组中的第 k 个 最小数组和。

## 解题思路

- 这一题是第 373 题的升级版。在第 373 题中，给定 2 个有序数组，要求分别从这 2 个数组中选出一个数组成一个数对，最终输出和最小的 K 组。这一题中给出的是 m*n 的矩阵。其实是将第 373 题的 2 个数组升级为了 m 个数组。无非外层多了一层循环。这层循环依次从每一行中选出一个数，先从第 0 行和第 1 行取数，找到前 K 小的组合以后，再从第 2 行取数，以此类推。其他做法和第 373 题一致。维护一个长度为 k 的最小堆。每次从堆中 pop 出最小的数组和 sum 和对应的下标 index，然后依次将下标向后移动一位，生成新的 sum，加入堆中。

## 代码

```go
package leetcode

import "container/heap"

func kthSmallest(mat [][]int, k int) int {
	if len(mat) == 0 || len(mat[0]) == 0 || k == 0 {
		return 0
	}
	prev := mat[0]
	for i := 1; i < len(mat); i++ {
		prev = kSmallestPairs(prev, mat[i], k)
	}
	if k < len(prev) {
		return -1
	}
	return prev[k-1]
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) []int {
	res := []int{}
	if len(nums2) == 0 {
		return res
	}
	pq := newPriorityQueue()
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(pq, &pddata{
			n1:    nums1[i],
			n2:    nums2[0],
			n2Idx: 0,
		})
	}
	for pq.Len() > 0 {
		i := heap.Pop(pq)
		data := i.(*pddata)
		res = append(res, data.n1+data.n2)
		k--
		if k <= 0 {
			break
		}
		idx := data.n2Idx
		idx++
		if idx >= len(nums2) {
			continue
		}
		heap.Push(pq, &pddata{
			n1:    data.n1,
			n2:    nums2[idx],
			n2Idx: idx,
		})
	}
	return res
}

type pddata struct {
	n1    int
	n2    int
	n2Idx int
}

type priorityQueue []*pddata

func newPriorityQueue() *priorityQueue {
	pq := priorityQueue([]*pddata{})
	heap.Init(&pq)
	return &pq
}

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].n1+pq[i].n2 < pq[j].n1+pq[j].n2 }
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	val := old[len(old)-1]
	old[len(old)-1] = nil
	*pq = old[0 : len(old)-1]
	return val
}

func (pq *priorityQueue) Push(i interface{}) {
	val := i.(*pddata)
	*pq = append(*pq, val)
}
```