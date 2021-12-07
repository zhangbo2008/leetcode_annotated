# [1647. Minimum Deletions to Make Character Frequencies Unique](https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/)


## 题目

A string `s` is called **good** if there are no two different characters in `s` that have the same **frequency**.

Given a string `s`, return *the **minimum** number of characters you need to delete to make* `s` ***good**.*

The **frequency** of a character in a string is the number of times it appears in the string. For example, in the string `"aab"`, the **frequency** of `'a'` is `2`, while the **frequency** of `'b'` is `1`.

**Example 1:**

```
Input: s = "aab"
Output: 0
Explanation: s is already good.

```

**Example 2:**

```
Input: s = "aaabbbcc"
Output: 2
Explanation: You can delete two 'b's resulting in the good string "aaabcc".
Another way it to delete one 'b' and one 'c' resulting in the good string "aaabbc".
```

**Example 3:**

```
Input: s = "ceabaacb"
Output: 2
Explanation: You can delete both 'c's resulting in the good string "eabaab".
Note that we only care about characters that are still in the string at the end (i.e. frequency of 0 is ignored).

```

**Constraints:**

- `1 <= s.length <= 105`
- `s` contains only lowercase English letters.

## 题目大意

如果字符串 s 中 不存在 两个不同字符 频次 相同的情况，就称 s 是 优质字符串 。

给你一个字符串 s，返回使 s 成为优质字符串需要删除的最小字符数。

字符串中字符的 频次 是该字符在字符串中的出现次数。例如，在字符串 "aab" 中，'a' 的频次是 2，而 'b' 的频次是 1 。

**提示：**

- `1 <= s.length <= 105`
- `s` 仅含小写英文字母

## 解题思路

- 给出一个字符串 s，要求输出使 s 变成“优质字符串”需要删除的最小字符数。“优质字符串”的定义是：字符串 s 中不存在频次相同的两个不同字符。
- 首先将 26 个字母在字符串中的频次分别统计出来，然后把频次从大到小排列，从频次大的开始，依次调整：例如，假设前一个和后一个频次相等，就把前一个字符删除一个，频次减一，再次排序，如果频次还相等，继续调整，如果频次不同了，游标往后移，继续调整后面的频次。直到所有的频次都不同了，就可以输出最终结果了。
- 这里需要注意频次为 0 的情况，即字母都被删光了。频次为 0 以后，就不需要再比较了。

## 代码

```go
package leetcode

import (
	"sort"
)

func minDeletions(s string) int {
	frequency, res := make([]int, 26), 0
	for i := 0; i < len(s); i++ {
		frequency[s[i]-'a']++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(frequency)))
	for i := 1; i <= 25; i++ {
		if frequency[i] == frequency[i-1] && frequency[i] != 0 {
			res++
			frequency[i]--
			sort.Sort(sort.Reverse(sort.IntSlice(frequency)))
			i--
		}
	}
	return res
}
```