# [1641. Count Sorted Vowel Strings](https://leetcode.com/problems/count-sorted-vowel-strings/)


## 题目

Given an integer `n`, return *the number of strings of length* `n` *that consist only of vowels (*`a`*,* `e`*,* `i`*,* `o`*,* `u`*) and are **lexicographically sorted**.*

A string `s` is **lexicographically sorted** if for all valid `i`, `s[i]` is the same as or comes before `s[i+1]` in the alphabet.

**Example 1:**

```
Input: n = 1
Output: 5
Explanation: The 5 sorted strings that consist of vowels only are ["a","e","i","o","u"].
```

**Example 2:**

```
Input: n = 2
Output: 15
Explanation: The 15 sorted strings that consist of vowels only are
["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
Note that "ea" is not a valid string since 'e' comes after 'a' in the alphabet.

```

**Example 3:**

```
Input: n = 33
Output: 66045

```

**Constraints:**

- `1 <= n <= 50`

## 题目大意

给你一个整数 n，请返回长度为 n 、仅由元音 (a, e, i, o, u) 组成且按 字典序排列 的字符串数量。

字符串 s 按 字典序排列 需要满足：对于所有有效的 i，s[i] 在字母表中的位置总是与 s[i+1] 相同或在 s[i+1] 之前。

## 解题思路

- 题目给的数据量并不大，第一个思路是利用 DFS 遍历打表法。时间复杂度 O(1)，空间复杂度 O(1)。
- 第二个思路是利用数学中的组合公式计算结果。题目等价于假设现在有 n 个字母，要求取 4 次球（可以选择不取）将字母分为 5 堆，问有几种取法。确定了取法以后，`a`，`e`，`i`，`o`，`u`，每个字母的个数就确定了，据题意要求按照字母序排序，那么最终字符串也就确定了。现在关注解决这个组合问题就可以了。把问题再转化一次，等价于，有 n+4 个字母，取 4 次，问有几种取法。+4 代表 4 个空操作，取走它们意味着不取。根据组合的数学定义，答案为 C(n+4,4)。

## 代码

```go
package leetcode

// 解法一 打表
func countVowelStrings(n int) int {
	res := []int{1, 5, 15, 35, 70, 126, 210, 330, 495, 715, 1001, 1365, 1820, 2380, 3060, 3876, 4845, 5985, 7315, 8855, 10626, 12650, 14950, 17550, 20475, 23751, 27405, 31465, 35960, 40920, 46376, 52360, 58905, 66045, 73815, 82251, 91390, 101270, 111930, 123410, 135751, 148995, 163185, 178365, 194580, 211876, 230300, 249900, 270725, 292825, 316251}
	return res[n]
}

func makeTable() []int {
	res, array := 0, []int{}
	for i := 0; i < 51; i++ {
		countVowelStringsDFS(i, 0, []string{}, []string{"a", "e", "i", "o", "u"}, &res)
		array = append(array, res)
		res = 0
	}
	return array
}

func countVowelStringsDFS(n, index int, cur []string, vowels []string, res *int) {
	vowels = vowels[index:]
	if len(cur) == n {
		(*res)++
		return
	}
	for i := 0; i < len(vowels); i++ {
		cur = append(cur, vowels[i])
		countVowelStringsDFS(n, i, cur, vowels, res)
		cur = cur[:len(cur)-1]
	}
}

// 解法二 数学方法 —— 组合
func countVowelStrings1(n int) int {
	return (n + 1) * (n + 2) * (n + 3) * (n + 4) / 24
}
```