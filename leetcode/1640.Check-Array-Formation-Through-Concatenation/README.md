# [1640. Check Array Formation Through Concatenation](https://leetcode.com/problems/check-array-formation-through-concatenation/)


## 题目

You are given an array of **distinct** integers `arr` and an array of integer arrays `pieces`, where the integers in `pieces` are **distinct**. Your goal is to form `arr` by concatenating the arrays in `pieces` **in any order**. However, you are **not** allowed to reorder the integers in each array `pieces[i]`.

Return `true` *if it is possible to form the array* `arr` *from* `pieces`. Otherwise, return `false`.

**Example 1:**

```
Input: arr = [85], pieces = [[85]]
Output: true
```

**Example 2:**

```
Input: arr = [15,88], pieces = [[88],[15]]
Output: true
Explanation: Concatenate [15] then [88]
```

**Example 3:**

```
Input: arr = [49,18,16], pieces = [[16,18,49]]
Output: false
Explanation: Even though the numbers match, we cannot reorder pieces[0].
```

**Example 4:**

```
Input: arr = [91,4,64,78], pieces = [[78],[4,64],[91]]
Output: true
Explanation: Concatenate [91] then [4,64] then [78]
```

**Example 5:**

```
Input: arr = [1,3,5,7], pieces = [[2,4,6,8]]
Output: false

```

**Constraints:**

- `1 <= pieces.length <= arr.length <= 100`
- `sum(pieces[i].length) == arr.length`
- `1 <= pieces[i].length <= arr.length`
- `1 <= arr[i], pieces[i][j] <= 100`
- The integers in `arr` are **distinct**.
- The integers in `pieces` are **distinct** (i.e., If we flatten pieces in a 1D array, all the integers in this array are distinct).

## 题目大意

给你一个整数数组 arr ，数组中的每个整数 互不相同 。另有一个由整数数组构成的数组 pieces，其中的整数也 互不相同 。请你以 任意顺序 连接 pieces 中的数组以形成 arr 。但是，不允许 对每个数组 pieces[i] 中的整数重新排序。如果可以连接 pieces 中的数组形成 arr ，返回 true ；否则，返回 false 。

## 解题思路

- 简单题。题目保证了 arr 中的元素唯一，所以可以用 map 把每个元素的 index 存起来，方便查找。遍历 pieces 数组，在每个一维数组中判断元素顺序是否和原 arr 元素相对顺序一致。这个时候就用 map 查找，如果顺序是一一相连的，那么就是正确的。有一个顺序不是一一相连，或者出现了 arr 不存在的元素，都返回 false。

## 代码

```go
package leetcode

func canFormArray(arr []int, pieces [][]int) bool {
	arrMap := map[int]int{}
	for i, v := range arr {
		arrMap[v] = i
	}
	for i := 0; i < len(pieces); i++ {
		order := -1
		for j := 0; j < len(pieces[i]); j++ {
			if _, ok := arrMap[pieces[i][j]]; !ok {
				return false
			}
			if order == -1 {
				order = arrMap[pieces[i][j]]
			} else {
				if arrMap[pieces[i][j]] == order+1 {
					order = arrMap[pieces[i][j]]
				} else {
					return false
				}
			}
		}
	}
	return true
}
```