# [792. Number of Matching Subsequences](https://leetcode.com/problems/number-of-matching-subsequences/)


## 题目

Given a string `s` and an array of strings `words`, return *the number of* `words[i]` *that is a subsequence of* `s`.

A **subsequence** of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

- For example, `"ace"` is a subsequence of `"abcde"`.

**Example 1:**

```
Input: s = "abcde", words = ["a","bb","acd","ace"]
Output: 3
Explanation: There are three strings in words that are a subsequence of s: "a", "acd", "ace".
```

**Example 2:**

```
Input: s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
Output: 2
```

**Constraints:**

- `1 <= s.length <= 5 * 104`
- `1 <= words.length <= 5000`
- `1 <= words[i].length <= 50`
- `s` and `words[i]` consist of only lowercase English letters.

## 题目大意

给定字符串 S 和单词字典 words, 求 words[i] 中是 S 的子序列的单词个数。

## 解题思路

- 如果将 words 数组内的字符串每次都在源字符串 S 中匹配，这种暴力解法超时。超时原因是对字符串 S 遍历了多次。是否有更加高效的方法呢？
- 把 words 数组内字符串按照首字母，分到 26 个桶中。从头开始遍历一遍源字符串 S，每扫一个字母，命中 26 个桶中其中一个桶，修改这个桶中的字符串。例如：当前遍历到了 'o'，此时桶中存的数据是 'a' : ['amy','aop'], 'o': ['oqp','onwn']，那么调整 'o' 桶中的数据后，各桶的状态为，'a' : ['amy','aop'], 'q': ['qp'], 'n': ['nwn']。从头到尾扫完整个字符串 S，某个桶中的字符串被清空，说明该桶中的字符串都符合 S 的子序列。将符合子序列的字符串个数累加起来即为最终答案。

## 代码

```go
package leetcode

func numMatchingSubseq(s string, words []string) int {
	hash, res := make([][]string, 26), 0
	for _, w := range words {
		hash[int(w[0]-'a')] = append(hash[int(w[0]-'a')], w)
	}
	for _, c := range s {
		words := hash[int(byte(c)-'a')]
		hash[int(byte(c)-'a')] = []string{}
		for _, w := range words {
			if len(w) == 1 {
				res += 1
				continue
			}
			hash[int(w[1]-'a')] = append(hash[int(w[1]-'a')], w[1:])
		}
	}
	return res
}
```