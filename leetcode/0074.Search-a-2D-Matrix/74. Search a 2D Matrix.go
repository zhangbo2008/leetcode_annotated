package leetcode
//下面方法是用二维转一维来二分查找.
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, low, high := len(matrix[0]), 0, len(matrix[0])*len(matrix)-1 //等价于一维的索引
	//其中m是矩阵的行数, low是最低索引,high是最高索引
	for low <= high {
		mid := low + (high-low)>>1  //然后二分.二分之后计算这个一维索引对应的二维索引位置. 用除法和%即可.
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
