# [1665. Minimum Initial Energy to Finish Tasks](https://leetcode.com/problems/minimum-initial-energy-to-finish-tasks/)

## 题目

You are given an array `tasks` where `tasks[i] = [actuali, minimumi]`:

- `actuali` is the actual amount of energy you **spend to finish** the `ith` task.
- `minimumi` is the minimum amount of energy you **require to begin** the `ith` task.

For example, if the task is `[10, 12]` and your current energy is `11`, you cannot start this task. However, if your current energy is `13`, you can complete this task, and your energy will be `3` after finishing it.

You can finish the tasks in **any order** you like.

Return *the **minimum** initial amount of energy you will need* *to finish all the tasks*.

**Example 1:**

```
Input: tasks = [[1,2],[2,4],[4,8]]
Output: 8
Explanation:
Starting with 8 energy, we finish the tasks in the following order:
    - 3rd task. Now energy = 8 - 4 = 4.
    - 2nd task. Now energy = 4 - 2 = 2.
    - 1st task. Now energy = 2 - 1 = 1.
Notice that even though we have leftover energy, starting with 7 energy does not work because we cannot do the 3rd task.
```

**Example 2:**

```
Input: tasks = [[1,3],[2,4],[10,11],[10,12],[8,9]]
Output: 32
Explanation:
Starting with 32 energy, we finish the tasks in the following order:
    - 1st task. Now energy = 32 - 1 = 31.
    - 2nd task. Now energy = 31 - 2 = 29.
    - 3rd task. Now energy = 29 - 10 = 19.
    - 4th task. Now energy = 19 - 10 = 9.
    - 5th task. Now energy = 9 - 8 = 1.
```

**Example 3:**

```
Input: tasks = [[1,7],[2,8],[3,9],[4,10],[5,11],[6,12]]
Output: 27
Explanation:
Starting with 27 energy, we finish the tasks in the following order:
    - 5th task. Now energy = 27 - 5 = 22.
    - 2nd task. Now energy = 22 - 2 = 20.
    - 3rd task. Now energy = 20 - 3 = 17.
    - 1st task. Now energy = 17 - 1 = 16.
    - 4th task. Now energy = 16 - 4 = 12.
    - 6th task. Now energy = 12 - 6 = 6.

```

**Constraints:**

- `1 <= tasks.length <= 105`
- `1 <= actuali <= minimumi <= 104`

## 题目大意

给你一个任务数组 tasks ，其中 tasks[i] = [actuali, minimumi] ：

- actual i 是完成第 i 个任务 需要耗费 的实际能量。
- minimum i 是开始第 i 个任务前需要达到的最低能量。

比方说，如果任务为 [10, 12] 且你当前的能量为 11 ，那么你不能开始这个任务。如果你当前的能量为 13 ，你可以完成这个任务，且完成它后剩余能量为 3 。你可以按照 任意顺序 完成任务。请你返回完成所有任务的 最少 初始能量。

## 解题思路

- 给出一个 task 数组，每个元素代表一个任务，每个任务有实际消费能量值和开始这个任务需要的最低能量。要求输出能完成所有任务的最少初始能量。
- 这一题直觉是贪心。先将任务按照 `minimum - actual` 进行排序。先完成差值大的任务，那么接下来的能量能最大限度的满足接下来的任务。这样可能完成所有任务的可能性越大。循环任务数组的时候，保存当前能量在 `cur` 中，如果当前能量不够开启下一个任务，那么这个差值就是需要弥补的，这些能量就是最少初始能量中的，所以加上这些差值能量。如果当前能量可以开启下一个任务，那么就更新当前能量，减去实际消耗的能量以后，再继续循环。循环结束就能得到最少初始能量了。

## 代码

```go
package leetcode

import (
	"sort"
)

func minimumEffort(tasks [][]int) int {
	sort.Sort(Task(tasks))
	res, cur := 0, 0
	for _, t := range tasks {
		if t[1] > cur {
			res += t[1] - cur
			cur = t[1] - t[0]
		} else {
			cur -= t[0]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Task define
type Task [][]int

func (task Task) Len() int {
	return len(task)
}

func (task Task) Less(i, j int) bool {
	t1, t2 := task[i][1]-task[i][0], task[j][1]-task[j][0]
	if t1 != t2 {
		return t2 < t1
	}
	return task[j][1] < task[i][1]
}

func (task Task) Swap(i, j int) {
	t := task[i]
	task[i] = task[j]
	task[j] = t
}
```