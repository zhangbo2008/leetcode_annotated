package leetcode

type position struct {
	x int
	y int
}

func solveSudoku(board [][]byte) { //走迷宫适合用dfs,因为空间开的小.找到一个解就返回了.不要求全解.
	pos, find := []position{}, false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				pos = append(pos, position{x: i, y: j})// pos记录所有的空白位置.
			}
		}
	}
	putSudoku(&board, pos, 0, &find)
}


//下面是对空白位置进行添加数字.  //position的index为本次函数放入的位置.
func putSudoku(board *[][]byte, pos []position, index int, succ *bool) {
	if *succ == true {
		return
	}
	if index == len(pos) {
		*succ = true
		return
	}
	for i := 1; i < 10; i++ {  //先check是否当前还能放数字,并且没有成功.才能继续放数字.
		if checkSudoku(board, pos[index], i) && !*succ {
			(*board)[pos[index].x][pos[index].y] = byte(i) + '0'  //board放入这个数字.
			putSudoku(board, pos, index+1, succ)
			if *succ == true {
				return
			}
			(*board)[pos[index].x][pos[index].y] = '.'//走出来说明匹配失败.重置为空.
		}
	}
}

func checkSudoku(board *[][]byte, pos position, val int) bool {
	// 判断横行是否有重复数字
	for i := 0; i < len((*board)[0]); i++ {//扫描同行是否有val这个数字.
		if (*board)[pos.x][i] != '.' && int((*board)[pos.x][i]-'0') == val {
			return false
		}
	}
	// 判断竖行是否有重复数字
	for i := 0; i < len((*board)); i++ {
		if (*board)[i][pos.y] != '.' && int((*board)[i][pos.y]-'0') == val {
			return false
		}
	}
	// 判断九宫格是否有重复数字
	posx, posy := pos.x-pos.x%3, pos.y-pos.y%3
	for i := posx; i < posx+3; i++ {
		for j := posy; j < posy+3; j++ {
			if (*board)[i][j] != '.' && int((*board)[i][j]-'0') == val {
				return false
			}
		}
	}
	return true
}
