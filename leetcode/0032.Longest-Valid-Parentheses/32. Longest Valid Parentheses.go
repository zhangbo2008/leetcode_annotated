package leetcode

// 解法一 栈
func longestValidParentheses(s string) int {
	stack, res := []int{}, 0
	stack = append(stack, -1) //先放一个-1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]  //每一次栈弹出的时候我们进行处理.
			if len(stack) == 0 {
				stack = append(stack, i)//说明这时候多了一个反括号
			} else {
				res = max(res, i-stack[len(stack)-1]) //这个值才是连续的能拼成括号的数量.
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法二 双指针
func longestValidParentheses1(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}
