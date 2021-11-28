package leetcode1


func twoSum(nums []int, target int) []int {
	m := make(map[int]int)// m 是一个map 从int到int的.//因为答案只有确定的一组,所以答案肯定在扫man的时候,他的target-v已经遍历过.所以下面的方法是对的.
	for k, v := range nums {//如果字典里面存着target-v了,那么返回这2个索引idx和k组成的数组即可.
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}