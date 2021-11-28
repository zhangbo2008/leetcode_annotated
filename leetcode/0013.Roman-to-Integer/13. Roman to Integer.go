package leetcode

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	num, lastint, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char] //根据字典进行转化
		if num < lastint {// 罗马字符一个小的写在大的左边表示减法, 写在右边表示加法.
			total = total - num
		} else {
			total = total + num
		}
		lastint = num
	}
	return total
}
