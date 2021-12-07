# [667. Beautiful Arrangement II](https://leetcode.com/problems/beautiful-arrangement-ii/)


## 题目

Given two integers `n` and `k`, you need to construct a list which contains `n` different positive integers ranging from `1` to `n` and obeys the following requirement:Suppose this list is [a1, a2, a3, ... , an], then the list [|a1 - a2|, |a2 - a3|, |a3 - a4|, ... , |an-1 - an|] has exactly `k` distinct integers.

If there are multiple answers, print any of them.

**Example 1:**

```
Input: n = 3, k = 1
Output: [1, 2, 3]
Explanation: The [1, 2, 3] has three different positive integers ranging from 1 to 3, and the [1, 1] has exactly 1 distinct integer: 1.
```

**Example 2:**

```
Input: n = 3, k = 2
Output: [1, 3, 2]
Explanation: The [1, 3, 2] has three different positive integers ranging from 1 to 3, and the [2, 1] has exactly 2 distinct integers: 1 and 2.
```

**Note:**

1. The `n` and `k` are in the range 1 <= k < n <= 10^4.

## 题目大意

给定两个整数 n 和 k，你需要实现一个数组，这个数组包含从 1 到 n 的 n 个不同整数，同时满足以下条件：

- 如果这个数组是 [a1, a2, a3, ... , an] ，那么数组 [|a1 - a2|, |a2 - a3|, |a3 - a4|, ... , |an-1 - an|] 中应该有且仅有 k 个不同整数；.
- 如果存在多种答案，你只需实现并返回其中任意一种.

## 解题思路

- 先考虑 `k` 最大值的情况。如果把末尾的较大值依次插入到前面的较小值中，形成 `[1，n，2，n-1，3，n-2，……]`，这样排列 `k` 能取到最大值 `n-1` 。`k` 最小值的情况是 `[1，2，3，4，……，n]`，`k` 取到的最小值是 1。那么 `k` 在 `[1，n-1]` 之间取值，该怎么排列呢？先顺序排列 `[1，2，3，4，……，n-k-1]`，这里有 `n-k-1` 个数，可以形成唯一一种差值。剩下 `k+1` 个数，形成 `k-1` 种差值。
- 这又回到了 `k` 最大值的取法了。`k` 取最大值的情况是 `n` 个数，形成 `n-1` 个不同种的差值。现在 `k+1` 个数，需要形成 `k` 种不同的差值。两者是同一个问题。那么剩下 `k` 个数的排列方法是 `[n-k，n-k+1，…，n]`，这里有 `k` 个数，注意代码实现时，注意 `k` 的奇偶性，如果 `k` 是奇数，“对半穿插”以后，正好匹配完，如果 `k` 是偶数，对半处的数 `n-k+(k+1)/2`，最后还需要单独加入到排列中。
- 可能有读者会问了，前面生成了 1 种差值，后面这部分又生产了 `k` 种差值，加起来不是 `k + 1` 种差值了么？这种理解是错误的。后面这段最后 2 个数字是 `n-k+(k+1)/2-1` 和 `n-k+(k+1)/2`，它们两者的差值是 1，和第一段构造的排列差值是相同的，都是 1。所以第一段构造了 1 种差值，第二段虽然构造了 `k` 种，但是需要去掉两段重复的差值 1，所以最终差值种类还是 `1 + k - 1 = k` 种。

## 代码

```go
package leetcode

func constructArray(n int, k int) []int {
	res := []int{}
	for i := 0; i < n-k-1; i++ {
		res = append(res, i+1)
	}
	for i := n - k; i < n-k+(k+1)/2; i++ {
		res = append(res, i)
		res = append(res, 2*n-k-i)
	}
	if k%2 == 0 {
		res = append(res, n-k+(k+1)/2)
	}
	return res
}
```