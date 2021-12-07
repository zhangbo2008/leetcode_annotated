# [551. Student Attendance Record I](https://leetcode.com/problems/student-attendance-record-i/)

## 题目

You are given a string `s` representing an attendance record for a student where each character signifies whether the student was absent, late, or present on that day. The record only contains the following three characters:

- `'A'`: Absent.
- `'L'`: Late.
- `'P'`: Present.

The student is eligible for an attendance award if they meet **both** of the following criteria:

- The student was absent (`'A'`) for **strictly** fewer than 2 days **total**.
- The student was **never** late (`'L'`) for 3 or more **consecutive** days.

Return `true` *if the student is eligible for an attendance award, or* `false` *otherwise*.

**Example 1:**

```
Input: s = "PPALLP"
Output: true
Explanation: The student has fewer than 2 absences and was never late 3 or more consecutive days.

```

**Example 2:**

```
Input: s = "PPALLL"
Output: false
Explanation: The student was late 3 consecutive days in the last 3 days, so is not eligible for the award.

```

**Constraints:**

- `1 <= s.length <= 1000`
- `s[i]` is either `'A'`, `'L'`, or `'P'`.

## 题目大意

给你一个字符串 s 表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：

- 'A'：Absent，缺勤
- 'L'：Late，迟到
- 'P'：Present，到场

如果学生能够 同时 满足下面两个条件，则可以获得出勤奖励：

- 按 总出勤 计，学生缺勤（'A'）严格 少于两天。
- 学生 不会 存在 连续 3 天或 连续 3 天以上的迟到（'L'）记录。

如果学生可以获得出勤奖励，返回 true ；否则，返回 false 。

## 解题思路

- 遍历字符串 s 求出 'A' 的总数量和连续 'L' 的最大数量。
- 比较 'A' 的数量是否小于 2 并且 'L' 的连续最大数量是否小于 3。

## 代码

```go
package leetcode

func checkRecord(s string) bool {
	numsA, maxL, numsL := 0, 0, 0
	for _, v := range s {
		if v == 'L' {
			numsL++
		} else {
			if numsL > maxL {
				maxL = numsL
			}
			numsL = 0
			if v == 'A' {
				numsA++
			}
		}
	}
	if numsL > maxL {
		maxL = numsL
	}
	return numsA < 2 && maxL < 3
}
```