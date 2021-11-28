package leetcode

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {// 如果左边的小,就左边的缩,哪边短缩哪边. 因为高的那边缩了,肯定面积变小.
		//这总情况直接剪枝掉.
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}






// 证明一下


// 首先我们知道这个问题暴力方法是遍历区间的left和right, 让他俩都从0----->n
//从而复杂度是N**2
// 而我们上面的算法是每次收缩短的边,让他向内走一格.
//证明一下这个.
// 只需要证明每一次向内走一格.
// 举例证明. 比如 3,2,4,6,7
//我们上来left=0 index   right=4 index
//我们如果证明了left向右走一格就 等价于遍历玩全部left=0,right=任意里面面积最大的情况.那么就足够证明这个问题了.
//其实这个证明left向右走一格就 等价于遍历玩全部left=0,right=任意里面面积最大的情况. 这个子问题是显然的.
// 看这个例子, left=0. right=4.那么他俩组成的矩阵最高是3. 而left=0,right=任意.那么得到的矩阵最高也是3.宽度只能比这个小于等于.所以我们每次收缩短的边即可.如果相等就随便收缩即可.





