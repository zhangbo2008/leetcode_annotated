# [1383. Maximum Performance of a Team](https://leetcode.com/problems/maximum-performance-of-a-team/)

## 题目

You are given two integers `n` and `k` and two integer arrays `speed` and `efficiency` both of length `n`. There are `n` engineers numbered from `1` to `n`. `speed[i]` and `efficiency[i]` represent the speed and efficiency of the `ith` engineer respectively.

Choose **at most** `k` different engineers out of the `n` engineers to form a team with the maximum **performance**.

The performance of a team is the sum of their engineers' speeds multiplied by the minimum efficiency among their engineers.

Return *the maximum performance of this team*. Since the answer can be a huge number, return it **modulo** `109 + 7`.

**Example 1:**

```
Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 2
Output: 60
Explanation:
We have the maximum performance of the team by selecting engineer 2 (with speed=10 and efficiency=4) and engineer 5 (with speed=5 and efficiency=7). That is, performance = (10 + 5) * min(4, 7) = 60.
```

**Example 2:**

```
Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 3
Output: 68
Explanation:
This is the same example as the first but k = 3. We can select engineer 1, engineer 2 and engineer 5 to get the maximum performance of the team. That is, performance = (2 + 10 + 5) * min(5, 4, 7) = 68.
```

**Example 3:**

```
Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 4
Output: 72
```

**Constraints:**

- `1 <= <= k <= n <= 105`
- `speed.length == n`
- `efficiency.length == n`
- `1 <= speed[i] <= 105`
- `1 <= efficiency[i] <= 108`

## 题目大意

公司有编号为 1 到 n 的 n 个工程师，给你两个数组 speed 和 efficiency ，其中 speed[i] 和 efficiency[i] 分别代表第 i 位工程师的速度和效率。请你返回由最多 k 个工程师组成的 最大团队表现值 ，由于答案可能很大，请你返回结果对 10^9 + 7 取余后的结果。团队表现值 的定义为：一个团队中「所有工程师速度的和」乘以他们「效率值中的最小值」。

## 解题思路

- 题目要求返回最大团队表现值，表现值需要考虑速度的累加和，和效率的最小值。即使速度快，效率的最小值很小，总的表现值还是很小。先将效率从大到小排序。从效率高的工程师开始选起，遍历过程中维护一个大小为 k 的速度最小堆。每次遍历都计算一次团队最大表现值。扫描完成，最大团队表现值也筛选出来了。具体实现见下面的代码。

## 代码

```go
package leetcode

import (
	"container/heap"
	"sort"
)

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	indexes := make([]int, n)
	for i := range indexes {
		indexes[i] = i
	}
	sort.Slice(indexes, func(i, j int) bool {
		return efficiency[indexes[i]] > efficiency[indexes[j]]
	})
	ph := speedHeap{}
	heap.Init(&ph)
	speedSum := 0
	var max int64
	for _, index := range indexes {
		if ph.Len() == k {
			speedSum -= heap.Pop(&ph).(int)
		}
		speedSum += speed[index]
		heap.Push(&ph, speed[index])
		max = Max(max, int64(speedSum)*int64(efficiency[index]))
	}
	return int(max % (1e9 + 7))
}

type speedHeap []int

func (h speedHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h speedHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h speedHeap) Len() int            { return len(h) }
func (h *speedHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *speedHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:h.Len()-1]
	return res
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
```