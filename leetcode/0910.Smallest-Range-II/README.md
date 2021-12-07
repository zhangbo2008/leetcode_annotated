# [910. Smallest Range II](https://leetcode.com/problems/smallest-range-ii/)

## 题目

Given an array `A` of integers, for each integer `A[i]` we need to choose **either `x = -K` or `x = K`**, and add `x` to `A[i]` **(only once)**.

After this process, we have some array `B`.

Return the smallest possible difference between the maximum value of `B` and the minimum value of `B`.

**Example 1:**

```c
Input: A = [1], K = 0
Output: 0
Explanation: B = [1]
```

**Example 2:**

```c
Input: A = [0,10], K = 2
Output: 6
Explanation: B = [2,8]
```

**Example 3:**

```c
Input: A = [1,3,6], K = 3
Output: 3
Explanation: B = [4,6,3]
```

**Note:**

1. `1 <= A.length <= 10000`
2. `0 <= A[i] <= 10000`
3. `0 <= K <= 10000`

## 题目大意

给你一个整数数组 A，对于每个整数 A[i]，可以选择 x = -K 或是 x = K （K 总是非负整数），并将 x 加到 A[i] 中。在此过程之后，得到数组 B。返回 B 的最大值和 B 的最小值之间可能存在的最小差值。

## 解题思路

- 简单题。先排序，找出 A 数组中最大的差值。然后循环扫一遍数组，利用双指针，选择 x = -K 或是 x = K ，每次选择都更新一次最大值和最小值之间的最小差值。循环一次以后便可以找到满足题意的答案。

## 代码

```go
package leetcode

import "sort"

func smallestRangeII(A []int, K int) int {
	n := len(A)
	sort.Ints(A)
	res := A[n-1] - A[0]
	for i := 0; i < n-1; i++ {
		a, b := A[i], A[i+1]
		high := max(A[n-1]-K, a+K)
		low := min(A[0]+K, b-K)
		res = min(res, high-low)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```