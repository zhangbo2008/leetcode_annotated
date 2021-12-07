# [1573. Number of Ways to Split a String](https://leetcode.com/problems/number-of-ways-to-split-a-string/)


## 题目

Given a binary string `s` (a string consisting only of '0's and '1's), we can split `s` into 3 **non-empty** strings s1, s2, s3 (s1+ s2+ s3 = s).

Return the number of ways `s` can be split such that the number of characters '1' is the same in s1, s2, and s3.

Since the answer may be too large, return it modulo 10^9 + 7.

**Example 1:**

```
Input: s = "10101"
Output: 4
Explanation: There are four ways to split s in 3 parts where each part contain the same number of letters '1'.
"1|010|1"
"1|01|01"
"10|10|1"
"10|1|01"

```

**Example 2:**

```
Input: s = "1001"
Output: 0

```

**Example 3:**

```
Input: s = "0000"
Output: 3
Explanation: There are three ways to split s in 3 parts.
"0|0|00"
"0|00|0"
"00|0|0"

```

**Example 4:**

```
Input: s = "100100010100110"
Output: 12

```

**Constraints:**

- `3 <= s.length <= 10^5`
- `s[i]` is `'0'` or `'1'`.

## 题目大意

给你一个二进制串 s  （一个只包含 0 和 1 的字符串），我们可以将 s 分割成 3 个 非空 字符串 s1, s2, s3 （s1 + s2 + s3 = s）。请你返回分割 s 的方案数，满足 s1，s2 和 s3 中字符 '1' 的数目相同。由于答案可能很大，请将它对 10^9 + 7 取余后返回。

## 解题思路

- 这一题是考察的排列组合的知识。根据题意，如果 1 的个数不是 3 的倍数，直接返回 -1。如果字符串里面没有 1，那么切分的方案就是组合，在 n-1 个字母里面选出 2 个位置。利用组合的计算方法，组合数是 (n-1) * (n-2) / 2 。
- 剩下的是 3 的倍数的情况。在字符串中选 2 个位置隔成 3 段。从第一段最后一个 1 到第二段第一个 1 之间的 0 的个数为 m1，从第二段最后一个 1 到第三段第一个 1 之间的 0 的个数为 m2。利用乘法原理，方案数为 m1 * m2。

## 代码

```go
package leetcode

func numWays(s string) int {
	ones := 0
	for _, c := range s {
		if c == '1' {
			ones++
		}
	}
	if ones%3 != 0 {
		return 0
	}
	if ones == 0 {
		return (len(s) - 1) * (len(s) - 2) / 2 % 1000000007
	}
	N, a, b, c, d, count := ones/3, 0, 0, 0, 0, 0
	for i, letter := range s {
		if letter == '0' {
			continue
		}
		if letter == '1' {
			count++
		}
		if count == N {
			a = i
		}
		if count == N+1 {
			b = i
		}
		if count == 2*N {
			c = i
		}
		if count == 2*N+1 {
			d = i
		}
	}
	return (b - a) * (d - c) % 1000000007
}
```