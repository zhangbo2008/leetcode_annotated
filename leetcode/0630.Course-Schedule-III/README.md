# [630. Course Schedule III](https://leetcode.com/problems/course-schedule-iii/)

## 题目

There are `n` different online courses numbered from `1` to `n`. You are given an array `courses` where `courses[i] = [durationi, lastDayi]` indicate that the `ith` course should be taken **continuously** for `durationi` days and must be finished before or on `lastDayi`.

You will start on the `1st` day and you cannot take two or more courses simultaneously.

Return *the maximum number of courses that you can take*.

**Example 1:**

```
Input: courses = [[100,200],[200,1300],[1000,1250],[2000,3200]]
Output: 3
Explanation:
There are totally 4 courses, but you can take 3 courses at most:
First, take the 1st course, it costs 100 days so you will finish it on the 100th day, and ready to take the next course on the 101st day.
Second, take the 3rd course, it costs 1000 days so you will finish it on the 1100th day, and ready to take the next course on the 1101st day.
Third, take the 2nd course, it costs 200 days so you will finish it on the 1300th day.
The 4th course cannot be taken now, since you will finish it on the 3300th day, which exceeds the closed date.

```

**Example 2:**

```
Input: courses = [[1,2]]
Output: 1

```

**Example 3:**

```
Input: courses = [[3,2],[4,3]]
Output: 0

```

**Constraints:**

- `1 <= courses.length <= 104`
- `1 <= durationi, lastDayi <= 104`

## 题目大意

这里有 n 门不同的在线课程，他们按从 1 到 n 编号。每一门课程有一定的持续上课时间（课程时间）t 以及关闭时间第 d 天。一门课要持续学习 t 天直到第 d 天时要完成，你将会从第 1 天开始。给出 n 个在线课程用 (t, d) 对表示。你的任务是找出最多可以修几门课。

## 解题思路

- 一般选课，任务的题目会涉及排序 + 贪心。此题同样如此。最多修几门课，采用贪心的思路。先将课程结束时间从小到大排序，优先选择结束时间靠前的课程，这样留给后面课程的时间越多，便可以修更多的课。对排好序的课程从前往后选课，不断累积时间。如果选择修当前课程，但是会超时，这时改调整了。对于已经选择的课程，都加入到最大堆中，遇到需要调整时，比较当前待考虑的课程时长是否比(堆中)已经选择课中时长最长的课时长短，即堆顶的课程时长短，剔除 pop 它，再选择这门时长短的课，并加入最大堆中。并更新累积时间。一层循环扫完所有课程，最终最大堆中包含课程的数目便是最多可以修的课程数。

## 代码

```go
package leetcode

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	maxHeap, time := &Schedule{}, 0
	heap.Init(maxHeap)
	for _, c := range courses {
		if time+c[0] <= c[1] {
			time += c[0]
			heap.Push(maxHeap, c[0])
		} else if (*maxHeap).Len() > 0 && (*maxHeap)[0] > c[0] {
			time -= heap.Pop(maxHeap).(int) - c[0]
			heap.Push(maxHeap, c[0])
		}
	}
	return (*maxHeap).Len()
}

type Schedule []int

func (s Schedule) Len() int           { return len(s) }
func (s Schedule) Less(i, j int) bool { return s[i] > s[j] }
func (s Schedule) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s *Schedule) Pop() interface{} {
	n := len(*s)
	t := (*s)[n-1]
	*s = (*s)[:n-1]
	return t
}
func (s *Schedule) Push(x interface{}) {
	*s = append(*s, x.(int))
}
```