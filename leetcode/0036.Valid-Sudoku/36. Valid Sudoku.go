package leetcode

import "strconv"

// 解法一 暴力遍历，时间复杂度 O(n^3)
func isValidSudoku(board [][]byte) bool {
	// 判断行 row
	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			cellVal := board[i][j : j+1]
			if string(cellVal) != "." {
				index, _ := strconv.Atoi(string(cellVal))
				if index > 9 || index < 1 {
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}
	// 判断列 column
	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			cellVal := board[j][i]
			if string(cellVal) != "." {
				index, _ := strconv.Atoi(string(cellVal))
				if index > 9 || index < 1 {
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}
	// 判断 9宫格 3X3 cell
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp := [10]int{}
			for ii := i * 3; ii < i*3+3; ii++ {
				for jj := j * 3; jj < j*3+3; jj++ {
					cellVal := board[ii][jj]
					if string(cellVal) != "." {
						index, _ := strconv.Atoi(string(cellVal))
						if tmp[index] == 1 {
							return false
						}
						tmp[index] = 1
					}
				}
			}
		}
	}
	return true
}











// 解法二 添加缓存，时间复杂度 O(n^2)=======直接看方法2.
func isValidSudoku1(board [][]byte) bool {
	//初始化3个表. rowbuf[i][j]表示i行里面j这个数字是否存在.
	//col同理
	//boxbuf是九宫格是否符合.
	rowbuf, colbuf, boxbuf := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {//bool数组初始化的默认值是false
		rowbuf[i] = make([]bool, 9)
		colbuf[i] = make([]bool, 9)
		boxbuf[i] = make([]bool, 9)
	}
	// 遍历一次，添加缓存
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {// 检查不是空的地方
				num := board[r][c] - '0' - byte(1) //因为前面2个ascii所以是uint8, 对应0到255. 所以最后用byte==uint8转化1再相见. 表示数字.  r/3*3+c/3 这个计算方法是算出这个九宫格是第几个. 计算原理是r/3*3 表示行号对应的第一个九宫格索引. c/3表示列号对应的九宫格索引偏移量.加上即可.
				if rowbuf[r][num] || colbuf[c][num] || boxbuf[r/3*3+c/3][num] {
					return false
				}
				rowbuf[r][num] = true
				colbuf[c][num] = true
				boxbuf[r/3*3+c/3][num] = true // r,c 转换到box方格中
			}
		}
	}
	return true
}
