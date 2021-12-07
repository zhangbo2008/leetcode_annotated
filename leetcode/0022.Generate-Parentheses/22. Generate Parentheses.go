package leetcode

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	findGenerateParenthesis(n, n, "", &res)  // 第一个参数表示还剩余多少个左括号需要放, 第二个参数表示还剩余多少个右括号需要放, 第三个是放之前的字符串, res是全局变量用来保存答案的.
	return res
}

func findGenerateParenthesis(lindex, rindex int, str string, res *[]string) {
	if lindex == 0 && rindex == 0 { //如果都不剩余了.那么答案加入str即可+
		*res = append(*res, str)
		return
	}
	if lindex > 0 {  //如果还剩余左括号就可以放.左括号.
		findGenerateParenthesis(lindex-1, rindex, str+"(", res)
	}
	if rindex > 0 && lindex < rindex {  //放右括号,必须右括号还有剩余的数量并且剩余的数量比左括号大才行.
		findGenerateParenthesis(lindex, rindex-1, str+")", res)
	}
}
