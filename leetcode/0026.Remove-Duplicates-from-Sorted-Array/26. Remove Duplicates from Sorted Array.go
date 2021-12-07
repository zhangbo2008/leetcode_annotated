package leetcode

// 解法一
func removeDuplicates(nums []int) int { //双指针来处理.
// 每次找到下一个数字跟上一个数字的第二个重复位置进行赋值.
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0

//last是上一个数字的初始索引,finder是下一个数字的第一个索引
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}//找到finder的索引
		nums[last+1] = nums[finder]//赋值过来
		last++ // last
	}

	return last + 1
}

// 解法二
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	lastNum := nums[length-1]
	i := 0
	for i = 0; i < length-1; i++ {
		if nums[i] == lastNum {
			break
		}
		if nums[i+1] == nums[i] {
			removeElement1(nums, i+1, nums[i])
			// fmt.Printf("此时 num = %v length = %v\n", nums, length)
		}
	}
	return i + 1
}

func removeElement1(nums []int, start, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := start
	for i := start; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
	return j
}
