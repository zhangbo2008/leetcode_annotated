# [1657. Determine if Two Strings Are Close](https://leetcode.com/problems/determine-if-two-strings-are-close/)


## 题目

Two strings are considered **close** if you can attain one from the other using the following operations:

- Operation 1: Swap any two **existing** characters.
    - For example, `abcde -> aecdb`
- Operation 2: Transform **every** occurrence of one **existing** character into another **existing** character, and do the same with the other character.
    - For example, `aacabb -> bbcbaa` (all `a`'s turn into `b`'s, and all `b`'s turn into `a`'s)

You can use the operations on either string as many times as necessary.

Given two strings, `word1` and `word2`, return `true` *if* `word1` *and* `word2` *are **close**, and* `false` *otherwise.*

**Example 1:**

```
Input: word1 = "abc", word2 = "bca"
Output: true
Explanation: You can attain word2 from word1 in 2 operations.
Apply Operation 1: "abc" -> "acb"
Apply Operation 1: "acb" -> "bca"

```

**Example 2:**

```
Input: word1 = "a", word2 = "aa"
Output: false
Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.

```

**Example 3:**

```
Input: word1 = "cabbba", word2 = "abbccc"
Output: true
Explanation: You can attain word2 from word1 in 3 operations.
Apply Operation 1: "cabbba" -> "caabbb"
Apply Operation 2: "caabbb" -> "baaccc"
Apply Operation 2: "baaccc" -> "abbccc"

```

**Example 4:**

```
Input: word1 = "cabbba", word2 = "aabbss"
Output: false
Explanation: It is impossible to attain word2 from word1, or vice versa, in any amount of operations.

```

**Constraints:**

- `1 <= word1.length, word2.length <= 105`
- `word1` and `word2` contain only lowercase English letters.

## 题目大意

如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 接近 ：

- 操作 1：交换任意两个 现有 字符。例如，abcde -> aecdb
- 操作 2：将一个 现有 字符的每次出现转换为另一个 现有 字符，并对另一个字符执行相同的操作。例如，aacabb -> bbcbaa（所有 a 转化为 b ，而所有的 b 转换为 a ）

你可以根据需要对任意一个字符串多次使用这两种操作。给你两个字符串，word1 和 word2 。如果 word1 和 word2 接近 ，就返回 true ；否则，返回 false 。

## 解题思路

- 判断 2 个字符串是否“接近”。“接近”的定义是能否通过交换 2 个字符或者 2 个字母互换，从一个字符串变换成另外一个字符串，如果存在这样的变换，即是“接近”。
- 先统计 2 个字符串的 26 个字母的频次，如果频次有不相同的，直接返回 false。在频次相同的情况下，再从小到大排序，再次扫描判断频次是否相同。
- 注意几种特殊情况：频次相同，再判断字母交换是否合法存在，如果字母不存在，输出 false。例如测试文件中的 case 5 。出现频次个数相同，但是频次不同。例如测试文件中的 case 6 。

## 代码

```go
package leetcode

import (
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	freqCount1, freqCount2 := make([]int, 26), make([]int, 26)
	for _, c := range word1 {
		freqCount1[c-97]++
	}
	for _, c := range word2 {
		freqCount2[c-97]++
	}
	for i := 0; i < 26; i++ {
		if (freqCount1[i] == freqCount2[i]) ||
			(freqCount1[i] > 0 && freqCount2[i] > 0) {
			continue
		}
		return false
	}
	sort.Ints(freqCount1)
	sort.Ints(freqCount2)
	for i := 0; i < 26; i++ {
		if freqCount1[i] != freqCount2[i] {
			return false
		}
	}
	return true
}
```