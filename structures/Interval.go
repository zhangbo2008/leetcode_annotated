package structures

// Interval 提供区间表示         ============写一下注释命名的含义
//   2表示to的含义 Ints 后面的s表示slice 一维切片就写一个s, 二维就写2个s
// Intss表示2维切片.
type Interval struct {
	Start int
	End   int
}

// Interval2Ints 把 Interval 转换成 整型切片
func Interval2Ints(i Interval) []int {
	return []int{i.Start, i.End}
}

// IntervalSlice2Intss 把 []Interval 转换成 [][]int
func IntervalSlice2Intss(is []Interval) [][]int {
	res := make([][]int, 0, len(is)) // make 第一个参数是目前大小,第二个参数是capacity
	//边学leetcode边学go就当了.
	for i := range is {
		res = append(res, Interval2Ints(is[i]))
	}
	return res
}

// Intss2IntervalSlice 把 [][]int 转换成 []Interval
func Intss2IntervalSlice(intss [][]int) []Interval {
	res := make([]Interval, 0, len(intss))
	for _, ints := range intss {
		res = append(res, Interval{Start: ints[0], End: ints[1]})
	}
	return res
}

// QuickSort define==========给出区间切片的快排算法,就是一个切片里面每一个元素都是Interval
func QuickSort(a []Interval, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionSort(a, lo, hi) //获取分割点.并且保证p之前的都比a[lo]小,p之后的都比a[lo]大.
	QuickSort(a, lo, p-1)
	QuickSort(a, p+1, hi)
}

func partitionSort(a []Interval, lo, hi int) int {
	pivot := a[hi] //pivot是最后一位.....
	i := lo - 1 //i取值为-leetcode1, i从0开始尝试,看是不是pivot最终停的位置.所以i就是记录有多少个j是正序的.
	for j := lo; j < hi; j++ {
		if (a[j].Start < pivot.Start) || (a[j].Start == pivot.Start && a[j].End < pivot.End) {
			//快拍的排序原则是 开始位置小就一定小. 开始位置一样,时候结尾位置小的算作小.
// j比pivot严格小,才进行交换.
			i++ // 注意这里面先i++ 所以不会出现负数索引. -1已经加到0了.
			a[j], a[i] = a[i], a[j]
		}
	}//经过这个for循环. i之前的都比pivot小. i之后的都比pivot大.
	a[i+1], a[hi] = a[hi], a[i+1]//所以i跟pivot换一下即可. 让pivot生效.
	return i + 1 //最后索引是i+leetcode1

	// 模拟一下 比如321   那么3 发生逆序 3跟i进行交换. i初始是-leetcode1.所以不换
	// 2跟1也是逆序也不换.
}
