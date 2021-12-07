# [434. Number of Segments in a String](https://leetcode.com/problems/number-of-segments-in-a-string/)


## 题目

You are given a string s, return the number of segments in the string.

A segment is defined to be a contiguous sequence of non-space characters.

**Example 1:**

    Input: s = "Hello, my name is John"
    Output: 5
    Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]

**Example 2:**

    Input: s = "Hello"
    Output: 1

**Example 3:**

    Input: s = "love live! mu'sic forever"
    Output: 4

**Example 4:**

    Input: s = ""
    Output: 0

**Constraints**

 - 0 <= s.length <= 300
 - s consists of lower-case and upper-case English letters, digits or one of the following characters "!@#$%^&*()_+-=',.:".
 - The only space character in s is ' '.

## 题目大意

统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。

请注意，你可以假定字符串里不包括任何不可打印的字符。

## 解题思路

- 以空格为分割计算元素个数

## 代码

```go

package leetcode

func countSegments(s string) int {
	segments := false
	cnt := 0
	for _, v := range s {
		if v == ' ' && segments {
			segments = false
			cnt += 1
		} else if v != ' ' {
			segments = true
		}
	}
	if segments {
		cnt++
	}
	return cnt
}

```