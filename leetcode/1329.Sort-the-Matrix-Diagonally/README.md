# [1329. Sort the Matrix Diagonally](https://leetcode.com/problems/sort-the-matrix-diagonally/)


## 题目

A **matrix diagonal** is a diagonal line of cells starting from some cell in either the topmost row or leftmost column and going in the bottom-right direction until reaching the matrix's end. For example, the **matrix diagonal** starting from `mat[2][0]`, where `mat` is a `6 x 3` matrix, includes cells `mat[2][0]`, `mat[3][1]`, and `mat[4][2]`.

Given an `m x n` matrix `mat` of integers, sort each **matrix diagonal** in ascending order and return *the resulting matrix*.

**Example 1:**

![https://assets.leetcode.com/uploads/2020/01/21/1482_example_1_2.png](https://assets.leetcode.com/uploads/2020/01/21/1482_example_1_2.png)

```
Input: mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
Output: [[1,1,1,1],[1,2,2,2],[1,2,3,3]]
```

**Constraints:**

- `m == mat.length`
- `n == mat[i].length`
- `1 <= m, n <= 100`
- `1 <= mat[i][j] <= 100`

## 题目大意

给你一个 m * n 的整数矩阵 mat ，请你将同一条对角线上的元素（从左上到右下）按升序排序后，返回排好序的矩阵。

## 解题思路

- 这道题思路很简单。按照对角线，把每条对角线的元素读取出来放在数组中。这里可以利用 map 保存这些数组。再将这些数组排序。最后按照对角线还原矩阵即可。

## 代码

```go
package leetcode

func diagonalSort(mat [][]int) [][]int {
	m, n, diagonalsMap := len(mat), len(mat[0]), make(map[int][]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagonalsMap[i-j] = append(diagonalsMap[i-j], mat[i][j])
		}
	}
	for _, v := range diagonalsMap {
		sort.Ints(v)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = diagonalsMap[i-j][0]
			diagonalsMap[i-j] = diagonalsMap[i-j][1:]
		}
	}
	return mat
}
```