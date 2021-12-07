# [1018. Binary Prefix Divisible By 5](https://leetcode.com/problems/binary-prefix-divisible-by-5/)


## 题目

Given an array `A` of `0`s and `1`s, consider `N_i`: the i-th subarray from `A[0]` to `A[i]` interpreted as a binary number (from most-significant-bit to least-significant-bit.)

Return a list of booleans `answer`, where `answer[i]` is `true` if and only if `N_i` is divisible by 5.

**Example 1:**

```
Input: [0,1,1]
Output: [true,false,false]
Explanation: 
The input numbers in binary are 0, 01, 011; which are 0, 1, and 3 in base-10.  Only the first number is divisible by 5, so answer[0] is true.

```

**Example 2:**

```
Input: [1,1,1]
Output: [false,false,false]

```

**Example 3:**

```
Input: [0,1,1,1,1,1]
Output: [true,false,false,false,true,false]

```

**Example 4:**

```
Input: [1,1,1,0,1]
Output: [false,false,false,false,false]

```

**Note:**

1. `1 <= A.length <= 30000`
2. `A[i]` is `0` or `1`

## 题目大意

给定由若干 0 和 1 组成的数组 A。我们定义 N_i：从 A[0] 到 A[i] 的第 i 个子数组被解释为一个二进制数（从最高有效位到最低有效位）。返回布尔值列表 answer，只有当 N_i 可以被 5 整除时，答案 answer[i] 为 true，否则为 false。

## 解题思路

- 简单题。每扫描数组中的一个数字，累计转换成二进制数对 5 取余，如果余数为 0，则存入 true，否则存入 false。

## 代码

```go
package leetcode

func prefixesDivBy5(a []int) []bool {
	res, num := make([]bool, len(a)), 0
	for i, v := range a {
		num = (num<<1 | v) % 5
		res[i] = num == 0
	}
	return res
}
```