# [368. Largest Divisible Subset](https://leetcode.com/problems/largest-divisible-subset/)


## 题目

Given a set of **distinct** positive integers `nums`, return the largest subset `answer` such that every pair `(answer[i], answer[j])` of elements in this subset satisfies:

- `answer[i] % answer[j] == 0`, or
- `answer[j] % answer[i] == 0`

If there are multiple solutions, return any of them.

**Example 1:**

```
Input: nums = [1,2,3]
Output: [1,2]
Explanation: [1,3] is also accepted.

```

**Example 2:**

```
Input: nums = [1,2,4,8]
Output: [1,2,4,8]

```

**Constraints:**

- `1 <= nums.length <= 1000`
- `1 <= nums[i] <= 2 * 109`
- All the integers in `nums` are **unique**.

## 题目大意

给你一个由 无重复 正整数组成的集合 nums ，请你找出并返回其中最大的整除子集 answer ，子集中每一元素对 (answer[i], answer[j]) 都应当满足：

- answer[i] % answer[j] == 0 ，或
- answer[j] % answer[i] == 0

如果存在多个有效解子集，返回其中任何一个均可。

## 解题思路

- 根据题目数据规模 1000，可以估计此题大概率是动态规划，并且时间复杂度是 O(n^2)。先将集合排序，以某一个小的数作为基准，不断的选择能整除的数加入集合。按照这个思路考虑，此题和第 300 题经典的 LIS 解题思路一致。只不过 LIS 每次选择更大的数，此题除了选择更大的数，只不过多了一个判断，这个更大的数能否整除当前集合里面的所有元素。按照此法一定可以找出最大的集合。
- 剩下的问题是如何输出最大集合。这道题的集合具有重叠子集的性质，例如 [2,4,8,16] 这个集合，长度是 4，它一定包含长度为 3 的子集，因为从它里面随便取 3 个数形成的子集也满足元素相互能整除的条件。同理，它也一定包含长度为 2，长度为 1 的子集。由于有这个性质，可以利用 dp 数组里面的数据，输出最大集合。例如，[2,4,6,8,9,13,16,40]，由动态规划可以找到最大集合是 [2,4,8,16]。长度为 4 的找到了，再找长度为 3 的，[2,4,8]，[2,4,40]。在最大集合中，最大元素是 16，所以 [2,4,40] 这个集合排除，它的最大元素大于 16。选定 [2,4,8] 这个集合，此时最大元素是 8 。以此类推，筛选到最后，便可以输出 [16,8,4,2] 这个组最大集合的答案了。

## 代码

```go
package leetcode

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp, res := make([]int, len(nums)), []int{}
	for i := range dp {
		dp[i] = 1
	}
	maxSize, maxVal := 1, 1
	for i := 1; i < len(nums); i++ {
		for j, v := range nums[:i] {
			if nums[i]%v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}
	if maxSize == 1 {
		return []int{nums[0]}
	}
	for i := len(nums) - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return res
}
```