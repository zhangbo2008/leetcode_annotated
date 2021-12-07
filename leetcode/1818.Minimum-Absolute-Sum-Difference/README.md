# [1818. Minimum Absolute Sum Difference](https://leetcode.com/problems/minimum-absolute-sum-difference/)

## 题目

You are given two positive integer arrays `nums1` and `nums2`, both of length `n`.

The **absolute sum difference** of arrays `nums1` and `nums2` is defined as the **sum** of `|nums1[i] - nums2[i]|` for each `0 <= i < n` (**0-indexed**).

You can replace **at most one** element of `nums1` with **any** other element in `nums1` to **minimize** the absolute sum difference.

Return the *minimum absolute sum difference **after** replacing at most one ****element in the array `nums1`.* Since the answer may be large, return it **modulo** `109 + 7`.

`|x|` is defined as:

- `x` if `x >= 0`, or
- `x` if `x < 0`.

**Example 1:**

```
Input: nums1 = [1,7,5], nums2 = [2,3,5]
Output: 3
Explanation:There are two possible optimal solutions:
- Replace the second element with the first: [1,7,5] => [1,1,5], or
- Replace the second element with the third: [1,7,5] => [1,5,5].
Both will yield an absolute sum difference of|1-2| + (|1-3| or |5-3|) + |5-5| =3.

```

**Example 2:**

```
Input: nums1 = [2,4,6,8,10], nums2 = [2,4,6,8,10]
Output: 0
Explanation:nums1 is equal to nums2 so no replacement is needed. This will result in an
absolute sum difference of 0.

```

**Example 3:**

```
Input: nums1 = [1,10,4,4,2,7], nums2 = [9,3,5,1,7,4]
Output: 20
Explanation:Replace the first element with the second: [1,10,4,4,2,7] => [10,10,4,4,2,7].
This yields an absolute sum difference of|10-9| + |10-3| + |4-5| + |4-1| + |2-7| + |7-4| = 20
```

**Constraints:**

- `n == nums1.length`
- `n == nums2.length`
- `1 <= n <= 10^5`
- `1 <= nums1[i], nums2[i] <= 10^5`

## 题目大意

给你两个正整数数组 nums1 和 nums2 ，数组的长度都是 n 。数组 nums1 和 nums2 的 绝对差值和 定义为所有 |nums1[i] - nums2[i]|（0 <= i < n）的 总和（下标从 0 开始）。你可以选用 nums1 中的 任意一个 元素来替换 nums1 中的 至多 一个元素，以 最小化 绝对差值和。在替换数组 nums1 中最多一个元素 之后 ，返回最小绝对差值和。因为答案可能很大，所以需要对 10^9 + 7 取余 后返回。

## 解题思路

- 如果不改变任何元素，绝对差值和为

$$\sum \left | nums1[i] - nums2[i] \right |$$

- 如果改变一个元素后，那么绝对差值和为

$$\begin{aligned}&\sum \left | nums1[i] - nums2[i] \right | - \left ( \left | nums1[i] - nums2[i] \right | - \left | nums1[j] - nums2[i] \right |\right )\\= &\sum \left | nums1[i] - nums2[i] \right | - \Delta \end{aligned}$$

题目要求返回最小绝对差值和，即求 

$$\Delta $$

的最大值。暴力枚举 nums1 和 nums2 中两两差值，找到 maxdiff。

## 代码

```go
package leetcode

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	diff := 0
	maxDiff := 0
	for i, n2 := range nums2 {
		d := abs(nums1[i] - n2)
		diff += d
		if maxDiff < d {
			t := 100001
			for _, n1 := range nums1 {
				maxDiff = max(maxDiff, d-min(t, abs(n1-n2)))
			}
		}
	}
	return (diff - maxDiff) % (1e9 + 7)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```