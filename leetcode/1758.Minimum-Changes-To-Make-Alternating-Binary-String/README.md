# [1758. Minimum Changes To Make Alternating Binary String](https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/)


## 题目

You are given a string `s` consisting only of the characters `'0'` and `'1'`. In one operation, you can change any `'0'` to `'1'` or vice versa.

The string is called alternating if no two adjacent characters are equal. For example, the string `"010"` is alternating, while the string `"0100"` is not.

Return *the **minimum** number of operations needed to make* `s` *alternating*.

**Example 1:**

```
Input: s = "0100"
Output: 1
Explanation: If you change the last character to '1', s will be "0101", which is alternating.
```

**Example 2:**

```
Input: s = "10"
Output: 0
Explanation: s is already alternating.
```

**Example 3:**

```
Input: s = "1111"
Output: 2
Explanation: You need two operations to reach "0101" or "1010".
```

**Constraints:**

- `1 <= s.length <= 104`
- `s[i]` is either `'0'` or `'1'`.

## 题目大意

你将得到一个仅包含字符“ 0”和“ 1”的字符串 `s`。 在一项操作中，你可以将任何 `'0'` 更改为 `'1'`，反之亦然。 如果两个相邻字符都不相等，则该字符串称为交替字符串。 例如，字符串“ 010”是交替的，而字符串“ 0100”则不是。 返回使 `s` 交替所需的最小操作数。

## 解题思路

- 简单题。利用数组下标奇偶交替性来判断交替字符串。交替字符串有 2 种，一个是 `'01010101……'` 还有一个是 `'1010101010……'`，这两个只需要计算出一个即可，另外一个利用 `len(s) - res` 就是答案。

## 代码

```go
package leetcode

func minOperations(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if int(s[i]-'0') != i%2 {
			res++
		}
	}
	return min(res, len(s)-res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```