# [1679. Max Number of K-Sum Pairs](https://leetcode.com/problems/max-number-of-k-sum-pairs/)


## 题目

You are given an integer array `nums` and an integer `k`.

In one operation, you can pick two numbers from the array whose sum equals `k` and remove them from the array.

Return *the maximum number of operations you can perform on the array*.

**Example 1:**

```
Input: nums = [1,2,3,4], k = 5
Output: 2
Explanation: Starting with nums = [1,2,3,4]:
- Remove numbers 1 and 4, then nums = [2,3]
- Remove numbers 2 and 3, then nums = []
There are no more pairs that sum up to 5, hence a total of 2 operations.
```

**Example 2:**

```
Input: nums = [3,1,3,4,3], k = 6
Output: 1
Explanation: Starting with nums = [3,1,3,4,3]:
- Remove the first two 3's, then nums = [1,4,3]
There are no more pairs that sum up to 6, hence a total of 1 operation.
```

**Constraints:**

- `1 <= nums.length <= 105`
- `1 <= nums[i] <= 109`
- `1 <= k <= 109`

## 题目大意

给你一个整数数组 nums 和一个整数 k 。每一步操作中，你需要从数组中选出和为 k 的两个整数，并将它们移出数组。返回你可以对数组执行的最大操作数。

## 解题思路

- 读完题第一感觉这道题是 TWO SUM 题目的加强版。需要找到所有满足和是 k 的数对。先考虑能不能找到两个数都是 k/2 ，如果能找到多个这样的数，可以先移除他们。其次在利用 TWO SUM 的思路，找出和为 k 的数对。利用 TWO SUM 里面 map 的做法，时间复杂度 O(n)。

## 代码

```go
package leetcode

// 解法一 优化版
func maxOperations(nums []int, k int) int {
	counter, res := make(map[int]int), 0
	for _, n := range nums {
		counter[n]++
	}
	if (k & 1) == 0 {
		res += counter[k>>1] >> 1
		// 能够由 2 个相同的数构成 k 的组合已经都排除出去了，剩下的一个单独的也不能组成 k 了
		// 所以这里要把它的频次置为 0 。如果这里不置为 0，下面代码判断逻辑还需要考虑重复使用数字的情况
		counter[k>>1] = 0
	}
	for num, freq := range counter {
		if num <= k/2 {
			remain := k - num
			if counter[remain] < freq {
				res += counter[remain]
			} else {
				res += freq
			}
		}
	}
	return res
}

// 解法二
func maxOperations_(nums []int, k int) int {
	counter, res := make(map[int]int), 0
	for _, num := range nums {
		counter[num]++
		remain := k - num
		if num == remain {
			if counter[num] >= 2 {
				res++
				counter[num] -= 2
			}
		} else {
			if counter[remain] > 0 {
				res++
				counter[remain]--
				counter[num]--
			}
		}
	}
	return res
}
```