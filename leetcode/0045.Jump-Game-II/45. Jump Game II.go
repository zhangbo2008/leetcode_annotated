package leetcode

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	needChoose, canReach, step := 0, 0, 0
	for i, x := range nums {
		if i+x > canReach { //i+x表示当前位置跳完能到的位置. 贪心直接每次都走到最大的那.
			canReach = i + x
			if canReach >= len(nums)-1 {
				return step + 1
			}
		}
		if i == needChoose {//表示走到了上一步的最大值.
			needChoose = canReach//那么就走下一个目标.
			step++
		}
	}
	return step
}
