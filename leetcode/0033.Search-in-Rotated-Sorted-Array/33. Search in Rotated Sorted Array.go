package leetcode

func search33(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[low] { // mid在数值大的一部分区间里
			if nums[low] <= target && target < nums[mid] { //再跟target比较.这时候我们比较mid所在的区域,因为只有这个区域是升序的.二分只有在有序的数列里面才能使用.
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // mid在数值小的一部分区间里=====跟上面是对称的.
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {//下面情况必然就是等于左右边界了.
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}
