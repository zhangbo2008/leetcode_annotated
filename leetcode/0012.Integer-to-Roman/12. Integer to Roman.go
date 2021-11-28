package leetcode

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res, i := "", 0
	for num != 0 {
		for values[i] > num {//只要当前的num大于values,那么我们就symbol加一个符号, nums剪去对应的阿拉伯数字.
			i++
		}
		num -= values[i]
		res += symbols[i]
	}
	return res
}
