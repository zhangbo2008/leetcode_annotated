package leetcode

func myAtoi(s string) int {
	maxInt, signAllowed, whitespaceAllowed, sign, digits := int64(2<<30), true, true, 1, []int{}
	//whitespaceAllowed表示允许识别带空格的.所以设置true就是题目需要的. sign存储符号digits存储数字.
	for _, c := range s {
		if c == ' ' && whitespaceAllowed {
			continue      //删除前面的空格
		}
		if signAllowed {
			if c == '+' {//遇到符号之后, 后面的就不允许再有空格和符号了.
				signAllowed = false
				whitespaceAllowed = false
				continue
			} else if c == '-' {
				sign = -1
				signAllowed = false
				whitespaceAllowed = false
				continue
			}
		}
		//到非数字时候退出
		if c < '0' || c > '9' {
			break
		}
		whitespaceAllowed, signAllowed = false, false//走到这里说明上面continue已经走过了.是正常数字了.所以
		//不允许有空格和符号了.
		digits = append(digits, int(c-48)) // 0的ascii编码是48
	}







	var num, place int64
	place, num = 1, 0
	lastLeading0Index := -1
	for i, d := range digits {
		if d == 0 {
			lastLeading0Index = i  //连续0的最后一个索引
		} else {
			break
		}
	}
	if lastLeading0Index > -1 {
		digits = digits[lastLeading0Index+1:]  //所以这个地方就去掉了开始全是0的部分
	}
	var rtnMax int64
	if sign > 0 {//算界限
		rtnMax = maxInt - 1
	} else {
		rtnMax = maxInt
	}


	digitsCount := len(digits)
	for i := digitsCount - 1; i >= 0; i-- {
		num += int64(digits[i]) * place//place是当前数位的表示10多少次幂.
		place *= 10
		if digitsCount-i > 10 || num > rtnMax {
			return int(int64(sign) * rtnMax)
		}
	}
	num *= int64(sign)
	return int(num)
}
