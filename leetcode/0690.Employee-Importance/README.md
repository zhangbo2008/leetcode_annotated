# [690. Employee Importance](https://leetcode.com/problems/employee-importance/)

## 题目

You are given a data structure of employee information, which includes the employee's **unique id**, their **importance value** and their **direct** subordinates' id.

For example, employee 1 is the leader of employee 2, and employee 2 is the leader of employee 3. They have importance value 15, 10 and 5, respectively. Then employee 1 has a data structure like [1, 15, [2]], and employee 2 has [2, 10, [3]], and employee 3 has [3, 5, []]. Note that although employee 3 is also a subordinate of employee 1, the relationship is **not direct**.

Now given the employee information of a company, and an employee id, you need to return the total importance value of this employee and all their subordinates.

**Example 1:**

```
Input: [[1, 5, [2, 3]], [2, 3, []], [3, 3, []]], 1
Output: 11
Explanation:
Employee 1 has importance value 5, and he has two direct subordinates: employee 2 and employee 3. They both have importance value 3. So the total importance value of employee 1 is 5 + 3 + 3 = 11.
```

**Note:**

1. One employee has at most one **direct** leader and may have several subordinates.
2. The maximum number of employees won't exceed 2000.

## 题目大意

给定一个保存员工信息的数据结构，它包含了员工 唯一的 id ，重要度 和 直系下属的 id 。比如，员工 1 是员工 2 的领导，员工 2 是员工 3 的领导。他们相应的重要度为 15 , 10 , 5 。那么员工 1 的数据结构是 [1, 15, [2]] ，员工 2的 数据结构是 [2, 10, [3]] ，员工 3 的数据结构是 [3, 5, []] 。注意虽然员工 3 也是员工 1 的一个下属，但是由于 并不是直系 下属，因此没有体现在员工 1 的数据结构中。现在输入一个公司的所有员工信息，以及单个员工 id ，返回这个员工和他所有下属的重要度之和。

## 解题思路

- 简单题。根据题意，DFS 或者 BFS 搜索找到所求 id 下属所有员工，累加下属员工的重要度，最后再加上这个员工本身的重要度，即为所求。

## 代码

```go
package leetcode

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	m, queue, res := map[int]*Employee{}, []int{id}, 0
	for _, e := range employees {
		m[e.Id] = e
	}
	for len(queue) > 0 {
		e := m[queue[0]]
		queue = queue[1:]
		if e == nil {
			continue
		}
		res += e.Importance
		for _, i := range e.Subordinates {
			queue = append(queue, i)
		}
	}
	return res
}
```