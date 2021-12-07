package leetcode
//最简单方法是记数然后再重写.
func sortColors(nums []int) {
	zero, one := 0, 0
	for i, n := range nums {
		nums[i] = 2
		if n <= 1 {
			nums[one] = 1
			one++
		}
		if n == 0 {
			nums[zero] = 0
			zero++
		}
	}
}
