package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			// no carry
			return digits
		}
		// carry
		digits[i] = 0  //满10就进位.
	}
	// all carry
	digits[0] = 1 //说明最高位置还有进位.
	digits = append(digits, 0)//后面补一个0.
	return digits
}
