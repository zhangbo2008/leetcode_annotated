package leetcode

// 解法一
func nextPermutation(nums []int) {
	i, j := 0, 0
	for i = len(nums) - 2; i >= 0; i-- {  //判断是否已经是最大序列.最大序列是严格的降序排列
		if nums[i] < nums[i+1] {
			break
		}
	}//    最后一个顺序的是ans. 数据是最后的上升自序列


	//
	//索引是i,j     那么i,j之间的数据一定>=j的数据. 反证这个结论.如果i,j之间的数据小于j了.
	// 那么让这个数据跟j交换就得到了一个比i,j交换更优的next permutation.
	// 所以就证明了i+1索引的数据一定比i大. 反过来证明如果i+1的索引比i大.那么i就是要找的.因为他符合我们
	//交换的原则, i的后面的数据有比i数据大的.

	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(&nums, i, j)
	}
	reverse(&nums, i+1, len(nums)-1)//计算最小字典序.
}

func reverse(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}













// 解法二
// [2,(3),6,5,4,1] -> 2,(4),6,5,(3),1 -> 2,4, 1,3,5,6
func nextPermutation1(nums []int) {
	var n = len(nums)
	var pIdx = checkPermutationPossibility(nums)
	if pIdx == -1 {
		reverse(&nums, 0, n-1)
		return
	}

	var rp = len(nums) - 1
	// start from right most to leftward,find the first number which is larger than PIVOT
	for rp > 0 {
		if nums[rp] > nums[pIdx] {
			swap(&nums, pIdx, rp)
			break
		} else {
			rp--
		}
	}
	// Finally, Reverse all elements which are right from pivot
	reverse(&nums, pIdx+1, n-1)
}

// checkPermutationPossibility returns 1st occurrence Index where
// value is in decreasing order(from right to left)
// returns -1 if not found(it's already in its last permutation)
func checkPermutationPossibility(nums []int) (idx int) {
	// search right to left for 1st number(from right) that is not in increasing order
	var rp = len(nums) - 1
	for rp > 0 {
		if nums[rp-1] < nums[rp] {
			idx = rp - 1
			return idx
		}
		rp--
	}
	return -1
}
