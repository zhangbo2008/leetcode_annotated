# [1748. Sum of Unique Elements](https://leetcode.com/problems/sum-of-unique-elements/)


## 题目

You are given an integer array `nums`. The unique elements of an array are the elements that appear **exactly once** in the array.

Return *the **sum** of all the unique elements of* `nums`.

**Example 1:**

```
Input: nums = [1,2,3,2]
Output: 4
Explanation: The unique elements are [1,3], and the sum is 4.
```

**Example 2:**

```
Input: nums = [1,1,1,1,1]
Output: 0
Explanation: There are no unique elements, and the sum is 0.
```

**Example 3:**

```
Input: nums = [1,2,3,4,5]
Output: 15
Explanation: The unique elements are [1,2,3,4,5], and the sum is 15.
```

**Constraints:**

- `1 <= nums.length <= 100`
- `1 <= nums[i] <= 100`

## 题目大意

给你一个整数数组 `nums` 。数组中唯一元素是那些只出现 **恰好一次** 的元素。请你返回 `nums` 中唯一元素的 **和** 。

## 解题思路

- 简单题。利用 map 统计出每个元素出现的频次。再累加所有频次为 1 的元素，最后输出累加和即可。

## 代码

```go
package leetcode

func sumOfUnique(nums []int) int {
	freq, res := make(map[int]int), 0
	for _, v := range nums {
		if _, ok := freq[v]; !ok {
			freq[v] = 0
		}
		freq[v]++
	}
	for k, v := range freq {
		if v == 1 {
			res += k
		}
	}
	return res
}
```