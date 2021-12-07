package leetcode
// 用简单方法, 首先正序遍历算出每一个索引i 前面的max, 逆序算每一个索引后面的max, 再
//遍历算每一个数字他跟 min(leftmax,rightmax)差即可.
func trap(height []int) int {
	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}
	return res
}
