# [1551. Minimum Operations to Make Array Equal](https://leetcode.com/problems/minimum-operations-to-make-array-equal/)


## 题目

You have an array `arr` of length `n` where `arr[i] = (2 * i) + 1` for all valid values of `i` (i.e. `0 <= i < n`).

In one operation, you can select two indices `x` and `y` where `0 <= x, y < n` and subtract `1` from `arr[x]` and add `1` to `arr[y]` (i.e. perform `arr[x] -=1` and `arr[y] += 1`). The goal is to make all the elements of the array **equal**. It is **guaranteed** that all the elements of the array can be made equal using some operations.

Given an integer `n`, the length of the array. Return *the minimum number of operations* needed to make all the elements of arr equal.

**Example 1:**

```
Input: n = 3
Output: 2
Explanation: arr = [1, 3, 5]
First operation choose x = 2 and y = 0, this leads arr to be [2, 3, 4]
In the second operation choose x = 2 and y = 0 again, thus arr = [3, 3, 3].
```

**Example 2:**

```
Input: n = 6
Output: 9
```

**Constraints:**

- `1 <= n <= 10^4`

## 题目大意

存在一个长度为 n 的数组 arr ，其中 arr[i] = (2 * i) + 1 （ 0 <= i < n ）。一次操作中，你可以选出两个下标，记作 x 和 y （ 0 <= x, y < n ）并使 arr[x] 减去 1 、arr[y] 加上 1 （即 arr[x] -=1 且 arr[y] += 1 ）。最终的目标是使数组中的所有元素都 相等 。题目测试用例将会 保证 ：在执行若干步操作后，数组中的所有元素最终可以全部相等。给你一个整数 n，即数组的长度。请你返回使数组 arr 中所有元素相等所需的 最小操作数 。

## 解题思路

- 这一题是数学题。题目给定的操作并不会使数组中所有元素之和变化，最终让所有元素相等，那么数组中所有元素的平均值即为最后数组中每一个元素的值。最少操作数的策略应该是以平均数为中心，中心右边的数减小，对称的中心左边的数增大。由于原数组是等差数列，两两元素之间相差 2，利用数学方法可以算出操作数。
- 数组长度分为奇数和偶数分别讨论。如果数组长度为奇数，所需要的操作数是：

    $$\begin{aligned} &\quad 2 + 4 + \cdots + 2\cdot\left\lfloor\frac{n}{2}\right\rfloor \\ &= \frac{1}{2}\left\lfloor\frac{n}{2}\right\rfloor\left(2\cdot\left\lfloor\frac{n}{2}\right\rfloor + 2\right) \\ &= \left\lfloor\frac{n}{2}\right\rfloor \left(\left\lfloor\frac{n}{2}\right\rfloor + 1\right) \\ &= \frac{n-1}{2}\left(\frac{n-1}{2} + 1\right) \\ &= \frac{n-1}{2}\cdot\frac{n+1}{2} \\ &= \frac{n^2-1}{4} \\ &= \left\lfloor\frac{n^2}{4}\right\rfloor \end{aligned}$$

    数组长度是偶数，所需要的操作数是：

    $$\begin{aligned} &\quad 1 + 3 + \cdots + \left(2\cdot\left\lfloor\frac{n}{2}\right\rfloor - 1\right) \\ &= \frac{1}{2}\left\lfloor\frac{n}{2}\right\rfloor\left(2\cdot\left\lfloor\frac{n}{2}\right\rfloor - 1 + 1\right)\\ &= \left(\left\lfloor\frac{n}{2}\right\rfloor\right)^2 \\ &= \frac{n^2}{4} \end{aligned}$$

    综上所述，最小操作数是 n^2/4

## 代码

```go
package leetcode

func minOperations(n int) int {
	return n * n / 4
}
```