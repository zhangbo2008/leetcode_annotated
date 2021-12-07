# [1736. Latest Time by Replacing Hidden Digits](https://leetcode.com/problems/latest-time-by-replacing-hidden-digits/)


## 题目

You are given a string `time` in the form of `hh:mm`, where some of the digits in the string are hidden (represented by `?`).

The valid times are those inclusively between `00:00` and `23:59`.

Return *the latest valid time you can get from* `time` *by replacing the hidden* *digits*.

**Example 1:**

```
Input: time = "2?:?0"
Output: "23:50"
Explanation: The latest hour beginning with the digit '2' is 23 and the latest minute ending with the digit '0' is 50.
```

**Example 2:**

```
Input: time = "0?:3?"
Output: "09:39"
```

**Example 3:**

```
Input: time = "1?:22"
Output: "19:22"
```

**Constraints:**

- `time` is in the format `hh:mm`.
- It is guaranteed that you can produce a valid time from the given string.

## 题目大意

给你一个字符串 time ，格式为 hh:mm（小时：分钟），其中某几位数字被隐藏（用 ? 表示）。有效的时间为 00:00 到 23:59 之间的所有时间，包括 00:00 和 23:59 。替换 time 中隐藏的数字，返回你可以得到的最晚有效时间。

## 解题思路

- 简单题。根据题意，需要找到最晚的有效时间。枚举时间 4 个位置即可。如果第 3 个位置是 ？，那么它最晚时间是 5；如果第 4 个位置是 ？，那么它最晚时间是 9；如果第 2 个位置是 ？，那么它最晚时间是 9；如果第 1 个位置是 ？，根据第 2 个位置判断，如果第 2 个位置是大于 3 的数，那么第一个位置最晚时间是 1，如果第 2 个位置是小于 3 的数那么第一个位置最晚时间是 2 。按照上述规则即可还原最晚时间。

## 代码

```go
package leetcode

func maximumTime(time string) string {
	timeb := []byte(time)
	if timeb[3] == '?' {
		timeb[3] = '5'
	}
	if timeb[4] == '?' {
		timeb[4] = '9'
	}
	if timeb[0] == '?' {
		if int(timeb[1]-'0') > 3 && int(timeb[1]-'0') < 10 {
			timeb[0] = '1'
		} else {
			timeb[0] = '2'
		}
	}
	if timeb[1] == '?' {
		timeb[1] = '9'
	}
	if timeb[0] == '2' && timeb[1] == '9' {
		timeb[1] = '3'
	}
	return string(timeb)
}
```