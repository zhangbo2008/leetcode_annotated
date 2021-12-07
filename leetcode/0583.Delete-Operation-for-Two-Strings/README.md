# [583. Delete Operation for Two Strings](https://leetcode.com/problems/delete-operation-for-two-strings/)


## 题目

Given two strings `word1` and `word2`, return *the minimum number of **steps** required to make* `word1` *and* `word2` *the same*.

In one **step**, you can delete exactly one character in either string.

**Example 1:**

```
Input: word1 = "sea", word2 = "eat"
Output: 2
Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
```

**Example 2:**

```
Input: word1 = "leetcode", word2 = "etco"
Output: 4
```

**Constraints:**

- `1 <= word1.length, word2.length <= 500`
- `word1` and `word2` consist of only lowercase English letters.

## 题目大意

给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。

## 解题思路

- 从题目数据量级判断，此题一定是 O(n^2) 动态规划题。定义 `dp[i][j]` 表示 `word1[:i]` 与 `word2[:j]` 匹配所删除的最少步数。如果 `word1[:i-1]` 与 `word2[:j-1]` 匹配，那么 `dp[i][j] = dp[i-1][j-1]`。如果 `word1[:i-1]` 与 `word2[:j-1]` 不匹配，那么需要考虑删除一次，所以 `dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j])`。所以动态转移方程是：

    $$dp[i][j] = \left\{\begin{matrix}dp[i-1][j-1]&, word1[i-1] == word2[j-1]\\ 1 + min(dp[i][j-1], dp[i-1][j])&, word1[i-1] \neq word2[j-1]\\\end{matrix}\right.$$

    最终答案存储在 `dp[len(word1)][len(word2)]` 中。

## 代码

```go
package leetcode

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < len(word2)+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
```