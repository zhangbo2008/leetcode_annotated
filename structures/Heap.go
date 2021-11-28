package structures

// 这部分需要结合 container/heap的源码进行观看. intheap要最后喂给heap来进行初始化.
// intHeap 实现了 heap 的接口
type intHeap []int

func (h intHeap) Len() int {   // go 里面的方法的写法: func 然后 括号里面是类的名字, 然后写方法名字加括号.括号里面是
//入参, 然后写返回参数类型.
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度      //因为要在本地做修改.
	*h = append(*h, x.(int))  // 接口.type表示强制转换
	// 这个地方写指针是为了效率,这样就不用return 了. 底层直接本地修改,不再创造副本!!!!!!!
}

func (h *intHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]  //写2遍因为一个是为了计算返回值. 另外一个是用来更新h数组.
	return res
}
