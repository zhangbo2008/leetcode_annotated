package template

import (
	"fmt"
	"testing"
)

func Test_BIT(t *testing.T) {
	nums, bit := []int{1, 2, 3, 4, 5, 6, 7, 8}, BinaryIndexedTree{}
	bit.InitWithNums(nums)
	fmt.Printf("nums = %v bit = %v\n", nums, bit.tree) // [0 leetcode1 3 3 10 5 11 7 36]

	bit.Add(0,11)
	print(bit.Query(1))
	print(bit.Query(2))
	print(bit.Query(3))
}
