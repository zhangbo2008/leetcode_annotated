package leetcode

import (
	"fmt"
	"strconv"
)




//  n=3, k=3 时候答案是 213
//   直接按照公式算. 每一个数位长n的第一个位置固定那么他的排列数是n-1的阶乘
// 所以每次除以n-1阶乘取整,然后把mod迭代下去

//首先所有的数是1,2,3
// 第一个位置是: 3=1*2!+1所以第一个数字是2,第二个数字是13的第一个排列.所以是213.


//n=4, k=9 答案2314


// 9=1*3!+3 所以第一位是2
//剩余134, 3=2!+1 所以第二个位置是3
//1=1!所以第三个位置是14的第一个.







func getPermutation(n int, k int) string {
	if k == 0 {
		return ""
	}
	used, p, res := make([]bool, n), []int{}, ""
	findPermutation(n, 0, &k, p, &res, &used)
	return res
}

func findPermutation(n, index int, k *int, p []int, res *string, used *[]bool) {
	fmt.Printf("n = %v index = %v k = %v p = %v res = %v user = %v\n", n, index, *k, p, *res, *used)
	if index == n {
		*k--
		if *k == 0 {
			for _, v := range p {
				*res += strconv.Itoa(v + 1)
			}
		}
		return
	}
	for i := 0; i < n; i++ {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, i)
			findPermutation(n, index+1, k, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}
