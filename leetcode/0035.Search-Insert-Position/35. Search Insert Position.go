package leetcode

func searchInsert(nums []int, target int) int { //找到比target小一点的那个索引.
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] >= target {
			high = mid - 1
		} else {// 这个结束条件就是 因为mid是小于target的.所以mid+1如果大于等于就停止.
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			low = mid + 1
		}
	}
	return 0
}
