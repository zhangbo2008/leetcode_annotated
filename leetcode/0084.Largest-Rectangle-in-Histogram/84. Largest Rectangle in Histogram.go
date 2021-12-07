package leetcode
//首先我们选择单调栈,是最大栈还是最小栈. 最大栈表示的是
func largestRectangleArea(heights []int) int {
	maxArea := 0
	n := len(heights) + 2
	// Add a sentry at the beginning and the end  //并且头尾都是数字0.
	getHeight := func(i int) int {
		if i == 0 || n-1 == i {
			return 0
		}
		return heights[i-1]
	}
	st := make([]int, 0, n/2)
	for i := 0; i < n; i++ {
		for len(st) > 0 && getHeight(st[len(st)-1]) > getHeight(i) {
			//如果st顶端的大于现在这个.那么就pop掉.所以一直保持最大栈.
			// pop stack  单调栈都是pop的时候进行处理.
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			maxArea = max(maxArea, getHeight(idx)*(i-st[len(st)-1]-1)) //这样每一个弹出都计算了他跟之后能拼成的长方形面积大小.  //最后的0会把之前的数字也pop出来.比如数组里面是123最后加入0的话.那么123都会pop出来.都会计算他跟之后的面积值.保证了计算的全面.
		}
		// push stack
		st = append(st, i)
	}
	return maxArea
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
