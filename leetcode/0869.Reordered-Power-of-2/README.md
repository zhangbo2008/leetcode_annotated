# [869. Reordered Power of 2](https://leetcode.com/problems/reordered-power-of-2/)


## 题目

Starting with a positive integer `N`, we reorder the digits in any order (including the original order) such that the leading digit is not zero.

Return `true` if and only if we can do this in a way such that the resulting number is a power of 2.

**Example 1:**

```
Input:1
Output:true
```

**Example 2:**

```
Input:10
Output:false
```

**Example 3:**

```
Input:16
Output:true
```

**Example 4:**

```
Input:24
Output:false
```

**Example 5:**

```
Input:46
Output:true
```

**Note:**

1. `1 <= N <= 10^9`

## 题目大意

给定正整数 N ，我们按任何顺序（包括原始顺序）将数字重新排序，注意其前导数字不能为零。如果我们可以通过上述方式得到 2 的幂，返回 true；否则，返回 false。

## 解题思路

- 将整数每个位上的所有排列看成字符串，那么题目转换为判断这些字符串是否和 2 的幂的字符串是否一致。判断的方法有很多种，笔者这里判断借助了一个 `map`。两个不同排列的字符串要相等，所有字符出现的频次必定一样。利用一个 `map` 统计它们各自字符的频次，最终都一致，则判定这两个字符串是满足题意的。
- 此题数据量比较小，在 `[1,10^9]` 这个区间内，2 的幂只有 30 几个，所以最终要判断的字符串就是这 30 几个。笔者这里没有打表了，采用更加一般的做法。数据量更大，此解法代码也能通过。

## 代码

```go
package leetcode

import "fmt"

func reorderedPowerOf2(n int) bool {
	sample, i := fmt.Sprintf("%v", n), 1
	for len(fmt.Sprintf("%v", i)) <= len(sample) {
		t := fmt.Sprintf("%v", i)
		if len(t) == len(sample) && isSame(t, sample) {
			return true
		}
		i = i << 1
	}
	return false
}

func isSame(t, s string) bool {
	m := make(map[rune]int)
	for _, v := range t {
		m[v]++
	}
	for _, v := range s {
		m[v]--
		if m[v] < 0 {
			return false
		}
		if m[v] == 0 {
			delete(m, v)
		}
	}
	return len(m) == 0
}
```