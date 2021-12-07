# [1539. Kth Missing Positive Number](https://leetcode.com/problems/kth-missing-positive-number/)

## 题目

Given an array `arr` of positive integers sorted in a **strictly increasing order**, and an integer `k`.

*Find the* `kth` *positive integer that is missing from this array.*

**Example 1:**

```
Input: arr = [2,3,4,7,11], k = 5
Output: 9
Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.
```

**Example 2:**

```
Input: arr = [1,2,3,4], k = 2
Output: 6
Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.
```

**Constraints:**

- `1 <= arr.length <= 1000`
- `1 <= arr[i] <= 1000`
- `1 <= k <= 1000`
- `arr[i] < arr[j]` for `1 <= i < j <= arr.length`

## 题目大意

给你一个 **严格升序排列** 的正整数数组 `arr` 和一个整数 `k` 。请你找到这个数组里第 `k` 个缺失的正整数。

## 解题思路

- 简单题。用一个变量从 1 开始累加，依次比对数组中是否存在，不存在的话就把 k - -，直到 k 为 0 的时候即是要输出的值。特殊情况，missing positive 都在数组之外，如例子 2 。

## 代码

```go
package leetcode

func findKthPositive(arr []int, k int) int {
	positive, index := 1, 0
	for index < len(arr) {
		if arr[index] != positive {
			k--
		} else {
			index++
		}
		if k == 0 {
			break
		}
		positive++
	}
	if k != 0 {
		positive += k - 1
	}
	return positive
}
```