# [1662. Check If Two String Arrays are Equivalent](https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/)


## 题目

Given two string arrays `word1` and `word2`, return **`true` *if the two arrays **represent** the same string, and* `false` *otherwise.*

A string is **represented** by an array if the array elements concatenated **in order** forms the string.

**Example 1:**

```
Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
Output: true
Explanation:
word1 represents string "ab" + "c" -> "abc"
word2 represents string "a" + "bc" -> "abc"
The strings are the same, so return true.
```

**Example 2:**

```
Input: word1 = ["a", "cb"], word2 = ["ab", "c"]
Output: false
```

**Example 3:**

```
Input: word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
Output: true
```

**Constraints:**

- `1 <= word1.length, word2.length <= 103`
- `1 <= word1[i].length, word2[i].length <= 103`
- `1 <= sum(word1[i].length), sum(word2[i].length) <= 103`
- `word1[i]` and `word2[i]` consist of lowercase letters.

## 题目大意

给你两个字符串数组 word1 和 word2 。如果两个数组表示的字符串相同，返回 true ；否则，返回 false 。数组表示的字符串 是由数组中的所有元素 按顺序 连接形成的字符串。

## 解题思路

- 简单题，依次拼接 2 个数组内的字符串，然后比较 str1 和 str2 是否相同即可。

## 代码

```go
package leetcode

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	str1, str2 := "", ""
	for i := 0; i < len(word1); i++ {
		str1 += word1[i]
	}
	for i := 0; i < len(word2); i++ {
		str2 += word2[i]
	}
	return str1 == str2
}
```