# [1646. Get Maximum in Generated Array](https://leetcode.com/problems/get-maximum-in-generated-array/)


## 题目

You are given an integer `n`. An array `nums` of length `n + 1` is generated in the following way:

- `nums[0] = 0`
- `nums[1] = 1`
- `nums[2 * i] = nums[i]` when `2 <= 2 * i <= n`
- `nums[2 * i + 1] = nums[i] + nums[i + 1]` when `2 <= 2 * i + 1 <= n`

Return *****the **maximum** integer in the array* `nums`.

**Example 1:**

```
Input: n = 7
Output: 3
Explanation: According to the given rules:
  nums[0] = 0
  nums[1] = 1
  nums[(1 * 2) = 2] = nums[1] = 1
  nums[(1 * 2) + 1 = 3] = nums[1] + nums[2] = 1 + 1 = 2
  nums[(2 * 2) = 4] = nums[2] = 1
  nums[(2 * 2) + 1 = 5] = nums[2] + nums[3] = 1 + 2 = 3
  nums[(3 * 2) = 6] = nums[3] = 2
  nums[(3 * 2) + 1 = 7] = nums[3] + nums[4] = 2 + 1 = 3
Hence, nums = [0,1,1,2,1,3,2,3], and the maximum is 3.

```

**Example 2:**

```
Input: n = 2
Output: 1
Explanation: According to the given rules, the maximum between nums[0], nums[1], and nums[2] is 1.

```

**Example 3:**

```
Input: n = 3
Output: 2
Explanation: According to the given rules, the maximum between nums[0], nums[1], nums[2], and nums[3] is 2.

```

**Constraints:**

- `0 <= n <= 100`

## 题目大意

给你一个整数 n 。按下述规则生成一个长度为 n + 1 的数组 nums ：

- nums[0] = 0
- nums[1] = 1
- 当 2 <= 2 * i <= n 时，nums[2 * i] = nums[i]
- 当 2 <= 2 * i + 1 <= n 时，nums[2 * i + 1] = nums[i] + nums[i + 1]

返回生成数组 nums 中的 最大值。

## 解题思路

- 给出一个 n + 1 的数组，并按照生成规则生成这个数组，求出这个数组中的最大值。
- 简单题，按照题意生成数组，边生成边记录和更新最大值即可。
- 注意边界条件，当 n 为 0 的时候，数组里面只有一个元素 0 。

## 代码

```go
package leetcode

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	nums, max := make([]int, n+1), 0
	nums[0], nums[1] = 0, 1
	for i := 0; i <= n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if 2*i >= 2 && 2*i <= n {
			nums[2*i] = nums[i]
		}
		if 2*i+1 >= 2 && 2*i+1 <= n {
			nums[2*i+1] = nums[i] + nums[i+1]
		}
	}
	return max
}
```