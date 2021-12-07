# [1846. Maximum Element After Decreasing and Rearranging](https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/)


## 题目

You are given an array of positive integers `arr`. Perform some operations (possibly none) on `arr` so that it satisfies these conditions:

- The value of the **first** element in `arr` must be `1`.
- The absolute difference between any 2 adjacent elements must be **less than or equal to** `1`. In other words, `abs(arr[i] - arr[i - 1]) <= 1` for each `i` where `1 <= i < arr.length` (**0-indexed**). `abs(x)` is the absolute value of `x`.

There are 2 types of operations that you can perform any number of times:

- **Decrease** the value of any element of `arr` to a **smaller positive integer**.
- **Rearrange** the elements of `arr` to be in any order.

Return *the **maximum** possible value of an element in* `arr` *after performing the operations to satisfy the conditions*.

**Example 1:**

```
Input: arr = [2,2,1,2,1]
Output: 2
Explanation:
We can satisfy the conditions by rearrangingarr so it becomes[1,2,2,2,1].
The largest element inarr is 2.

```

**Example 2:**

```
Input: arr = [100,1,1000]
Output: 3
Explanation:
One possible way to satisfy the conditions is by doing the following:
1. Rearrangearr so it becomes[1,100,1000].
2. Decrease the value of the second element to 2.
3. Decrease the value of the third element to 3.
Nowarr = [1,2,3], whichsatisfies the conditions.
The largest element inarr is 3.
```

**Example 3:**

```
Input: arr = [1,2,3,4,5]
Output: 5
Explanation: The array already satisfies the conditions, and the largest element is 5.

```

**Constraints:**

- `1 <= arr.length <= 10^5`
- `1 <= arr[i] <= 10^9`

## 题目大意

给你一个正整数数组 arr 。请你对 arr 执行一些操作（也可以不进行任何操作），使得数组满足以下条件：

- arr 中 第一个 元素必须为 1 。
- 任意相邻两个元素的差的绝对值 小于等于 1 ，也就是说，对于任意的 1 <= i < arr.length （数组下标从 0 开始），都满足 abs(arr[i] - arr[i - 1]) <= 1 。abs(x) 为 x 的绝对值。

你可以执行以下 2 种操作任意次：

- 减小 arr 中任意元素的值，使其变为一个 更小的正整数 。
- 重新排列 arr 中的元素，你可以以任意顺序重新排列。

请你返回执行以上操作后，在满足前文所述的条件下，arr 中可能的 最大值 。

## 解题思路

- 正整数数组 arr 第一个元素必须为 1，且两两元素绝对值小于等于 1，那么 arr 最大值肯定不大于 n。采用贪心的策略，先统计所有元素出现的次数，大于 n 的元素出现次数都累加到 n 上。然后从 1 扫描到 n，遇到“空隙”（出现次数为 0 的元素），便将最近一个出现次数大于 1 的元素“挪”过来填补“空隙”。题目所求最大值出现在，“填补空隙”之后，数组从左往右连续的最右端。

## 代码

```go
package leetcode

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	count := make([]int, n+1)
	for _, v := range arr {
		count[min(v, n)]++
	}
	miss := 0
	for _, c := range count[1:] {
		if c == 0 {
			miss++
		} else {
			miss -= min(c-1, miss)
		}
	}
	return n - miss
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```