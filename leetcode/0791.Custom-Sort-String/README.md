# [791. Custom Sort String](https://leetcode.com/problems/custom-sort-string/)


## 题目

`order` and `str` are strings composed of lowercase letters. In `order`, no letter occurs more than once.

`order` was sorted in some custom order previously. We want to permute the characters of `str` so that they match the order that `order` was sorted. More specifically, if `x` occurs before `y` in `order`, then `x` should occur before `y` in the returned string.

Return any permutation of `str` (as a string) that satisfies this property.

```
Example:Input:
order = "cba"
str = "abcd"
Output: "cbad"
Explanation:
"a", "b", "c" appear in order, so the order of "a", "b", "c" should be "c", "b", and "a".
Since "d" does not appear in order, it can be at any position in the returned string. "dcba", "cdba", "cbda" are also valid outputs.

```

**Note:**

- `order` has length at most `26`, and no character is repeated in `order`.
- `str` has length at most `200`.
- `order` and `str` consist of lowercase letters only.

## 题目大意

字符串 S 和 T 只包含小写字符。在 S 中，所有字符只会出现一次。S 已经根据某种规则进行了排序。我们要根据 S 中的字符顺序对 T 进行排序。更具体地说，如果 S 中 x 在 y 之前出现，那么返回的字符串中 x 也应出现在 y 之前。返回任意一种符合条件的字符串 T。

## 解题思路

- 题目只要求 T 中包含 S 的字符串有序，所以可以先将 T 中包含 S 的字符串排好序，然后再拼接上其他字符。S 字符串最长为 26 位，先将 S 中字符的下标向左偏移 30，并将偏移后的下标值存入字典中。再把 T 字符串按照字典中下标值进行排序。S 中出现的字符对应的下标经过处理以后变成了负数，S 中未出现的字符的下标还是正数。所以经过排序以后，S 中出现的字符按照原有顺序排列在前面，S 中未出现的字符依次排在后面。

## 代码

```go
package leetcode

import "sort"

func customSortString(order string, str string) string {
	magic := map[byte]int{}
	for i := range order {
		magic[order[i]] = i - 30
	}
	byteSlice := []byte(str)
	sort.Slice(byteSlice, func(i, j int) bool {
		return magic[byteSlice[i]] < magic[byteSlice[j]]
	})
	return string(byteSlice)
}
```