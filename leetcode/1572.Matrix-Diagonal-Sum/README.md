# [1572. Matrix Diagonal Sum](https://leetcode.com/problems/matrix-diagonal-sum/)


## 题目

Given a square matrix `mat`, return the sum of the matrix diagonals.

Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.

**Example 1:**

![https://assets.leetcode.com/uploads/2020/08/14/sample_1911.png](https://assets.leetcode.com/uploads/2020/08/14/sample_1911.png)

```
Input: mat = [[1,2,3],
              [4,5,6],
              [7,8,9]]
Output: 25
Explanation:Diagonals sum: 1 + 5 + 9 + 3 + 7 = 25
Notice that element mat[1][1] = 5 is counted only once.

```

**Example 2:**

```
Input: mat = [[1,1,1,1],
              [1,1,1,1],
              [1,1,1,1],
              [1,1,1,1]]
Output: 8

```

**Example 3:**

```
Input: mat = [[5]]
Output: 5

```

**Constraints:**

- `n == mat.length == mat[i].length`
- `1 <= n <= 100`
- `1 <= mat[i][j] <= 100`

## 题目大意

给你一个正方形矩阵 mat，请你返回矩阵对角线元素的和。请你返回在矩阵主对角线上的元素和副对角线上且不在主对角线上元素的和。

## 解题思路

- 简单题。根据题意，把主对角线和副对角线上的元素相加。
- 如果正方形矩阵的长度 n 为奇数，相加的结果需要减去 mat[n/2][n/2]。

## 代码

```go
package leetcode

func diagonalSum(mat [][]int) int {
	n := len(mat)
	ans := 0
	for pi := 0; pi < n; pi++ {
		ans += mat[pi][pi]
	}
	for si, sj := n-1, 0; sj < n; si, sj = si-1, sj+1 {
		ans += mat[si][sj]
	}
	if n%2 == 0 {
		return ans
	}
	return ans - mat[n/2][n/2]
}
```