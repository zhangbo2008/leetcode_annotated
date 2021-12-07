package leetcode

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0//双指针, 一个i一个j.   其中i是负责扫描后面的数哪个要交换的位置
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {//找到要交换的位置.跟j交换.
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}
