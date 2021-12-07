# [417. Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/)


## 题目

Given an `m x n` matrix of non-negative integers representing the height of each unit cell in a continent, the "Pacific ocean" touches the left and top edges of the matrix and the "Atlantic ocean" touches the right and bottom edges.

Water can only flow in four directions (up, down, left, or right) from a cell to another one with height equal or lower.

Find the list of grid coordinates where water can flow to both the Pacific and Atlantic ocean.

**Note:**

1. The order of returned grid coordinates does not matter.
2. Both m and n are less than 150.

**Example:**

```
Given the following 5x5 matrix:

  Pacific ~   ~   ~   ~   ~
       ~  1   2   2   3  (5) *
       ~  3   2   3  (4) (4) *
       ~  2   4  (5)  3   1  *
       ~ (6) (7)  1   4   5  *
       ~ (5)  1   1   2   4  *
          *   *   *   *   * Atlantic

Return:

[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (positions with parentheses in above matrix).

```

## 题目大意

给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。

## 解题思路

- 暴力解法，利用 DFS 把二维数据按照行优先搜索一遍，分别标记出太平洋和大西洋水流能到达的位置。再按照列优先搜索一遍，标记出太平洋和大西洋水流能到达的位置。最后两者都能到达的坐标即为所求。

## 代码

```go
package leetcode

import "math"

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	row, col, res := len(matrix), len(matrix[0]), make([][]int, 0)
	pacific, atlantic := make([][]bool, row), make([][]bool, row)
	for i := 0; i < row; i++ {
		pacific[i] = make([]bool, col)
		atlantic[i] = make([]bool, col)
	}
	for i := 0; i < row; i++ {
		dfs(matrix, i, 0, &pacific, math.MinInt32)
		dfs(matrix, i, col-1, &atlantic, math.MinInt32)
	}
	for j := 0; j < col; j++ {
		dfs(matrix, 0, j, &pacific, math.MinInt32)
		dfs(matrix, row-1, j, &atlantic, math.MinInt32)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if atlantic[i][j] && pacific[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func dfs(matrix [][]int, row, col int, visited *[][]bool, height int) {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		return
	}
	if (*visited)[row][col] || matrix[row][col] < height {
		return
	}
	(*visited)[row][col] = true
	dfs(matrix, row+1, col, visited, matrix[row][col])
	dfs(matrix, row-1, col, visited, matrix[row][col])
	dfs(matrix, row, col+1, visited, matrix[row][col])
	dfs(matrix, row, col-1, visited, matrix[row][col])
}
```