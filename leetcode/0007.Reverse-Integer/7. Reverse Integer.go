package leetcode

func reverse7(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10  //每次%10就能得到他的最低位.  然后每次之前的数*10加上最低位即可.
		x = x / 10           //  /10就能去掉最低位.
	}
	//再判断是否越界
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}
