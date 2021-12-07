package leetcode

import "strings"

// 解法一 ==========正确方法是kmp才对.复杂度O(m+n)
func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}

// 解法二      //这个太狗了自动忽略
func strStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
