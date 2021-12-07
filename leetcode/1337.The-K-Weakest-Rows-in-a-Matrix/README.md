# [1337. The K Weakest Rows in a Matrix](https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/)


## 题目

Given a `m * n` matrix `mat` of *ones* (representing soldiers) and *zeros* (representing civilians), return the indexes of the `k` weakest rows in the matrix ordered from the weakest to the strongest.

A row ***i*** is weaker than row ***j***, if the number of soldiers in row ***i*** is less than the number of soldiers in row ***j***, or they have the same number of soldiers but ***i*** is less than ***j***. Soldiers are **always** stand in the frontier of a row, that is, always *ones* may appear first and then *zeros*.

**Example 1:**

```
Input: mat = 
[[1,1,0,0,0],
 [1,1,1,1,0],
 [1,0,0,0,0],
 [1,1,0,0,0],
 [1,1,1,1,1]], 
k = 3
Output: [2,0,3]
Explanation: 
The number of soldiers for each row is: 
row 0 -> 2 
row 1 -> 4 
row 2 -> 1 
row 3 -> 2 
row 4 -> 5 
Rows ordered from the weakest to the strongest are [2,0,3,1,4]

```

**Example 2:**

```
Input: mat = 
[[1,0,0,0],
 [1,1,1,1],
 [1,0,0,0],
 [1,0,0,0]], 
k = 2
Output: [0,2]
Explanation: 
The number of soldiers for each row is: 
row 0 -> 1 
row 1 -> 4 
row 2 -> 1 
row 3 -> 1 
Rows ordered from the weakest to the strongest are [0,2,3,1]

```

**Constraints:**

- `m == mat.length`
- `n == mat[i].length`
- `2 <= n, m <= 100`
- `1 <= k <= m`
- `matrix[i][j]` is either 0 **or** 1.

## 题目大意

给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。

## 解题思路

- 简单题。第一个能想到的解题思路是，先统计每一行 1 的个数，然后将结果进行排序，按照 1 的个数从小到大排序，如果 1 的个数相同，再按照行号从小到大排序。排好序的数组取出前 K 位即为答案。
- 此题还有第二种解法。在第一种解法中，并没有用到题目中“军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。”这一条件。由于有了这个条件，使得如果按照列去遍历，最先出现 0 的行，则是最弱的行。行号小的先被遍历到，所以相同数量 1 的行，行号小的会排在前面。最后记得再添加上全 1 的行。同样，最终输出取出前 K 位即为答案。此题解法二才是最优雅最高效的解法。

## 代码

```go
package leetcode

func kWeakestRows(mat [][]int, k int) []int {
	res := []int{}
	for j := 0; j < len(mat[0]); j++ {
		for i := 0; i < len(mat); i++ {
			if mat[i][j] == 0 && ((j == 0) || (mat[i][j-1] != 0)) {
				res = append(res, i)
			}
		}
	}
	for i := 0; i < len(mat); i++ {
		if mat[i][len(mat[0])-1] == 1 {
			res = append(res, i)
		}
	}
	return res[:k]
}
```