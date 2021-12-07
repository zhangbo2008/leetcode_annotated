# [1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit](https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/)


## 题目

Given an array of integers `nums` and an integer `limit`, return the size of the longest **non-empty** subarray such that the absolute difference between any two elements of this subarray is less than or equal to `limit`*.*

**Example 1:**

```
Input: nums = [8,2,4,7], limit = 4
Output: 2 
Explanation: All subarrays are: 
[8] with maximum absolute diff |8-8| = 0 <= 4.
[8,2] with maximum absolute diff |8-2| = 6 > 4. 
[8,2,4] with maximum absolute diff |8-2| = 6 > 4.
[8,2,4,7] with maximum absolute diff |8-2| = 6 > 4.
[2] with maximum absolute diff |2-2| = 0 <= 4.
[2,4] with maximum absolute diff |2-4| = 2 <= 4.
[2,4,7] with maximum absolute diff |2-7| = 5 > 4.
[4] with maximum absolute diff |4-4| = 0 <= 4.
[4,7] with maximum absolute diff |4-7| = 3 <= 4.
[7] with maximum absolute diff |7-7| = 0 <= 4. 
Therefore, the size of the longest subarray is 2.
```

**Example 2:**

```
Input: nums = [10,1,2,4,7,2], limit = 5
Output: 4 
Explanation: The subarray [2,4,7,2] is the longest since the maximum absolute diff is |2-7| = 5 <= 5.
```

**Example 3:**

```
Input: nums = [4,2,2,2,4,4,2,2], limit = 0
Output: 3
```

**Constraints:**

- `1 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^9`
- `0 <= limit <= 10^9`

## 题目大意

给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。如果不存在满足条件的子数组，则返回 0 。

## 解题思路

- 最开始想到的思路是利用滑动窗口遍历一遍数组，每个窗口内排序，取出最大最小值。滑动窗口遍历一次的时间复杂度是 O(n)，所以此题时间复杂度是否高效落在了排序算法上了。由于前后 2 个窗口数据是有关联的，仅仅只变动了 2 个数据（左窗口移出的数据和右窗口移进的数据），所以排序没有必要每次都重新排序。这里利用二叉排序树来排序，添加和删除元素时间复杂度是 O(log n)，这种方法总的时间复杂度是 O(n log n)。空间复杂度 O(n)。
- 二叉排序树的思路是否还有再优化的空间？答案是有。二叉排序树内维护了所有结点的有序关系，但是这个关系是多余的。此题只需要找到最大值和最小值，并不需要除此以外节点的有序信息。所以用二叉排序树是大材小用了。可以换成 2 个单调队列，一个维护窗口内的最大值，另一个维护窗口内的最小值。这样优化以后，时间复杂度降低到 O(n)，空间复杂度 O(n)。具体实现见代码。
- 单调栈的题还有第 42 题，第 84 题，第 496 题，第 503 题，第 739 题，第 856 题，第 901 题，第 907 题，第 1130 题，第 1425 题，第 1673 题。

## 代码

```go
package leetcode

func longestSubarray(nums []int, limit int) int {
	minStack, maxStack, left, res := []int{}, []int{}, 0, 0
	for right, num := range nums {
		for len(minStack) > 0 && nums[minStack[len(minStack)-1]] > num {
			minStack = minStack[:len(minStack)-1]
		}
		minStack = append(minStack, right)
		for len(maxStack) > 0 && nums[maxStack[len(maxStack)-1]] < num {
			maxStack = maxStack[:len(maxStack)-1]
		}
		maxStack = append(maxStack, right)
		if len(minStack) > 0 && len(maxStack) > 0 && nums[maxStack[0]]-nums[minStack[0]] > limit {
			if left == minStack[0] {
				minStack = minStack[1:]
			}
			if left == maxStack[0] {
				maxStack = maxStack[1:]
			}
			left++
		}
		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}
```