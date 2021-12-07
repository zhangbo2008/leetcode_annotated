# [1689. Partitioning Into Minimum Number Of Deci-Binary Numbers](https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/)

## 题目

A decimal number is called **deci-binary** if each of its digits is either `0` or `1` without any leading zeros. For example, `101` and `1100` are **deci-binary**, while `112` and `3001` are not.

Given a string `n` that represents a positive decimal integer, return *the **minimum** number of positive **deci-binary** numbers needed so that they sum up to* `n`*.*

**Example 1:**

```
Input: n = "32"
Output: 3
Explanation: 10 + 11 + 11 = 32
```

**Example 2:**

```
Input: n = "82734"
Output: 8
```

**Example 3:**

```
Input: n = "27346209830709182346"
Output: 9
```

**Constraints:**

- `1 <= n.length <= 105`
- `n` consists of only digits.
- `n` does not contain any leading zeros and represents a positive integer.

## 题目大意

如果一个十进制数字不含任何前导零，且每一位上的数字不是 0 就是 1 ，那么该数字就是一个 十-二进制数 。例如，101 和 1100 都是 十-二进制数，而 112 和 3001 不是。给你一个表示十进制整数的字符串 n ，返回和为 n 的 十-二进制数 的最少数目。

## 解题思路

- 这一题也算是简单题，相通了以后，代码就 3 行。
- 要想由 01 组成的十进制数组成 n，只需要在 n 这个数的各个数位上依次排上 0 和 1 即可。例如 n = 23423723，这是一个 8 位数。最大数字是 7，所以至少需要 7 个数累加能得到这个 n。这 7 个数的百位都为 1，其他数位按需求取 0 和 1 即可。例如万位是 2，那么这 7 个数中任找 2 个数的万位是 1 ，其他 5 个数的万位是 0 即可。

## 代码

```go
package leetcode

func minPartitions(n string) int {
	res := 0
	for i := 0; i < len(n); i++ {
		if int(n[i]-'0') > res {
			res = int(n[i] - '0')
		}
	}
	return res
}
```