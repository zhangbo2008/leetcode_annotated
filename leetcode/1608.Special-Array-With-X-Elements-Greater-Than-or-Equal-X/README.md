# [1608. Special Array With X Elements Greater Than or Equal X](https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/)

## 题目

You are given an array `nums` of non-negative integers. `nums` is considered **special** if there exists a number `x` such that there are **exactly** `x` numbers in `nums` that are **greater than or equal to** `x`.

Notice that `x` **does not** have to be an element in `nums`.

Return `x` *if the array is **special**, otherwise, return* `-1`. It can be proven that if `nums` is special, the value for `x` is **unique**.

**Example 1:**

```
Input: nums = [3,5]
Output: 2
Explanation: There are 2 values (3 and 5) that are greater than or equal to 2.
```

**Example 2:**

```
Input: nums = [0,0]
Output: -1
Explanation: No numbers fit the criteria for x.
If x = 0, there should be 0 numbers >= x, but there are 2.
If x = 1, there should be 1 number >= x, but there are 0.
If x = 2, there should be 2 numbers >= x, but there are 0.
x cannot be greater since there are only 2 numbers in nums.
```

**Example 3:**

```
Input: nums = [0,4,3,0,4]
Output: 3
Explanation: There are 3 values that are greater than or equal to 3.
```

**Example 4:**

```
Input: nums = [3,6,7,7,0]
Output: -1
```

**Constraints:**

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 1000`

## 题目大意

给你一个非负整数数组 nums 。如果存在一个数 x ，使得 nums 中恰好有 x 个元素 大于或者等于 x ，那么就称 nums 是一个 特殊数组 ，而 x 是该数组的 特征值 。（注意： x 不必 是 nums 的中的元素。）如果数组 nums 是一个 特殊数组 ，请返回它的特征值 x 。否则，返回 -1 。可以证明的是，如果 nums 是特殊数组，那么其特征值 x 是 唯一的 。

## 解题思路

- 简单题。抓住题干中给的证明，特征值是唯一的。先将数组从小到大排序，下标的含义与特征值就等价了。下标 `i` 代表大于等于 `nums[i]` 的元素有 `len(nums) - i` 个，那么从第 0 个下标的元素开始遍历，如果这个元素都大于 `len(nums)`，那么后面 `len(nums)` 个元素也都大于等于它，特征值就找到了。如果特征值减一以后，仍然满足 `nums[i] >= x`，说明满足条件的值有多个，这一点不满足特征值唯一性，可以直接返回 -1 了。下标继续右移，特征值继续减一。如果最终循环结束依旧找不到特征值，返回 -1 。

## 代码

```go
package leetcode

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	x := len(nums)
	for _, num := range nums {
		if num >= x {
			return x
		}
		x--
		if num >= x {
			return -1
		}
	}
	return -1
}
```