# [1684. Count the Number of Consistent Strings](https://leetcode.com/problems/count-the-number-of-consistent-strings/)


## 题目

You are given a string `allowed` consisting of **distinct** characters and an array of strings `words`. A string is **consistent** if all characters in the string appear in the string `allowed`.

Return *the number of **consistent** strings in the array* `words`.

**Example 1:**

```
Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
Output: 2
Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.

```

**Example 2:**

```
Input: allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
Output: 7
Explanation: All strings are consistent.

```

**Example 3:**

```
Input: allowed = "cad", words = ["cc","acd","b","ba","bac","bad","ac","d"]
Output: 4
Explanation: Strings "cc", "acd", "ac", and "d" are consistent.

```

**Constraints:**

- `1 <= words.length <= 104`
- `1 <= allowed.length <= 26`
- `1 <= words[i].length <= 10`
- The characters in `allowed` are **distinct**.
- `words[i]` and `allowed` contain only lowercase English letters.

## 题目大意

给你一个由不同字符组成的字符串 `allowed` 和一个字符串数组 `words` 。如果一个字符串的每一个字符都在 `allowed` 中，就称这个字符串是 一致字符串 。

请你返回 `words` 数组中 一致字符串 的数目。

## 解题思路

- 简单题。先将 `allowed` 转化成 map。将 `words` 数组中每个单词的字符都在 map 中查找一遍，如果都存在就累加 res。如果有不存在的字母，不累加。最终输出 res 即可。

## 代码

```go
package leetcode

func countConsistentStrings(allowed string, words []string) int {
	allowedMap, res, flag := map[rune]int{}, 0, true
	for _, str := range allowed {
		allowedMap[str]++
	}
	for i := 0; i < len(words); i++ {
		flag = true
		for j := 0; j < len(words[i]); j++ {
			if _, ok := allowedMap[rune(words[i][j])]; !ok {
				flag = false
				break
			}
		}
		if flag {
			res++
		}
	}
	return res
}
```