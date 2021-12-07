# [1048. Longest String Chain](https://leetcode.com/problems/longest-string-chain/)


## 题目

Given a list of words, each word consists of English lowercase letters.

Let's say `word1` is a predecessor of `word2` if and only if we can add exactly one letter anywhere in `word1` to make it equal to `word2`. For example, `"abc"` is a predecessor of `"abac"`.

A *word chain* is a sequence of words `[word_1, word_2, ..., word_k]` with `k >= 1`, where `word_1` is a predecessor of `word_2`, `word_2` is a predecessor of `word_3`, and so on.

Return the longest possible length of a word chain with words chosen from the given list of `words`.

**Example 1:**

```
Input: words = ["a","b","ba","bca","bda","bdca"]
Output: 4
Explanation: One of the longest word chain is "a","ba","bda","bdca".
```

**Example 2:**

```
Input: words = ["xbc","pcxbcf","xb","cxbc","pcxbc"]
Output: 5
```

**Constraints:**

- `1 <= words.length <= 1000`
- `1 <= words[i].length <= 16`
- `words[i]` only consists of English lowercase letters.

## 题目大意

给出一个单词列表，其中每个单词都由小写英文字母组成。如果我们可以在 word1 的任何地方添加一个字母使其变成 word2，那么我们认为 word1 是 word2 的前身。例如，"abc" 是 "abac" 的前身。词链是单词 [word_1, word_2, ..., word_k] 组成的序列，k >= 1，其中 word_1 是 word_2 的前身，word_2 是 word_3 的前身，依此类推。从给定单词列表 words 中选择单词组成词链，返回词链的最长可能长度。

## 解题思路

- 从这题的数据规模上分析，可以猜出此题是 DFS 或者 DP 的题。简单暴力的方法是以每个字符串为链条的起点开始枚举之后的字符串，两两判断能否构成满足题意的前身字符串。这种做法包含很多重叠子问题，例如 a 和 b 能构成前身字符串，以 c 为起点的字符串链条可能用到 a 和 b，以 d 为起点的字符串链条也可能用到 a 和 b。顺其自然，考虑用 DP 的思路解题。
- 先将 words 字符串数组排序，然后用 poss 数组记录下每种长度字符串的在排序数组中的起始下标。然后逆序往前递推。因为初始条件只能得到以最长字符串为起始的字符串链长度为 1 。每选择一个起始字符串，从它的长度 + 1 的每个字符串 j 开始比较，是否能为其前身字符串。如果能构成前身字符串，那么 dp[i] = max(dp[i], 1+dp[j])。最终递推到下标为 0 的字符串。最终输出整个递推过程中的最大长度即为所求。

## 代码

```go
package leetcode

import "sort"

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	poss, res := make([]int, 16+2), 0
	for i, w := range words {
		if poss[len(w)] == 0 {
			poss[len(w)] = i
		}
	}
	dp := make([]int, len(words))
	for i := len(words) - 1; i >= 0; i-- {
		dp[i] = 1
		for j := poss[len(words[i])+1]; j < len(words) && len(words[j]) == len(words[i])+1; j++ {
			if isPredecessor(words[j], words[i]) {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isPredecessor(long, short string) bool {
	i, j := 0, 0
	wasMismatch := false
	for j < len(short) {
		if long[i] != short[j] {
			if wasMismatch {
				return false
			}
			wasMismatch = true
			i++
			continue
		}
		i++
		j++
	}
	return true
}
```