package leetcode

import "strconv"

// 解法一
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	sz := len(arr)
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

// 解法二 数字转字符串          =======推荐用这个方法,偷懒不容易错.
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)//转化为字符串
	length := len(s)
	for i := 0; i <= length/2; i++ {//然后双指针从两边向中间走即可.
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
