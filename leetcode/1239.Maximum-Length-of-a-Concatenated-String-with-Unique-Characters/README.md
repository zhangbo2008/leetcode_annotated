# [1239. Maximum Length of a Concatenated String with Unique Characters](https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/)

## 题目

Given an array of strings `arr`. String `s` is a concatenation of a sub-sequence of `arr` which have **unique characters**.

Return *the maximum possible length* of `s`.

**Example 1:**

```
Input: arr = ["un","iq","ue"]
Output: 4
Explanation: All possible concatenations are "","un","iq","ue","uniq" and "ique".
Maximum length is 4.
```

**Example 2:**

```
Input: arr = ["cha","r","act","ers"]
Output: 6
Explanation: Possible solutions are "chaers" and "acters".
```

**Example 3:**

```
Input: arr = ["abcdefghijklmnopqrstuvwxyz"]
Output: 26
```

**Constraints:**

- `1 <= arr.length <= 16`
- `1 <= arr[i].length <= 26`
- `arr[i]` contains only lower case English letters.

## 题目大意

给定一个字符串数组 arr，字符串 s 是将 arr 某一子序列字符串连接所得的字符串，如果 s 中的每一个字符都只出现过一次，那么它就是一个可行解。请返回所有可行解 s 中最长长度。

## 解题思路

- 每个字符串数组可以想象为 26 位的 0101 二进制串。出现的字符对应的位上标记为 1，没有出现的字符对应的位上标记为 0 。如果一个字符串中包含重复的字符，那么它所有 1 的个数一定不等于字符串的长度。如果 2 个字符串每个字母都只出现了一次，那么它们俩对应的二进制串 mask 相互与运算的结果一定为 0 ，即 0，1 互补了。利用这个特点，深搜所有解，保存出最长可行解的长度即可。

## 代码

```go
package leetcode

import (
	"math/bits"
)

func maxLength(arr []string) int {
	c, res := []uint32{}, 0
	for _, s := range arr {
		var mask uint32
		for _, c := range s {
			mask = mask | 1<<(c-'a')
		}
		if len(s) != bits.OnesCount32(mask) { // 如果字符串本身带有重复的字符，需要排除
			continue
		}
		c = append(c, mask)
	}
	dfs(c, 0, 0, &res)
	return res
}

func dfs(c []uint32, index int, mask uint32, res *int) {
	*res = max(*res, bits.OnesCount32(mask))
	for i := index; i < len(c); i++ {
		if mask&c[i] == 0 {
			dfs(c, i+1, mask|c[i], res)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```