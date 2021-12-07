package leetcode

func isValid(s string) bool {
	// 空字符串直接返回 true
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0) //栈里面纯ascii码也就是char数据, go里面用rune表示char数据.
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
		} else if ((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {


			stack = stack[:len(stack)-1]    //右括号时候需要判断是否符合弹出条件.首先如果v是右大括号,那么栈不能为空,并且栈顶是左大括号.同理3个条件.如果成功.那么出战即可.
		} else {
			return false
		}
	}
	return len(stack) == 0
}
