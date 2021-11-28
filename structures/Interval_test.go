package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Interval2Ints(t *testing.T) {
	ast := assert.New(t)

	actual := Interval2Ints(Interval{Start: 1, End: 2})
	expected := []int{1, 2}
	ast.Equal(expected, actual)
}

func Test_IntervalSlice2Intss(t *testing.T) {
	ast := assert.New(t)

	actual := IntervalSlice2Intss(
		[]Interval{
			{
				Start: 1,
				End:   2,
			},
			{
				Start: 3,
				End:   4,
			},
		},
	)
	expected := [][]int{
		{1, 2},
		{3, 4},
	}

	ast.Equal(expected, actual)
}
func Test_Intss2IntervalSlice(t *testing.T) {
	ast := assert.New(t)

	expected := []Interval{
		{
			Start: 1,
			End:   2,
		},
		{
			Start: 3,
			End:   4,
		},
	}
	actual := Intss2IntervalSlice([][]int{
		{1, 2},
		{3, 4},
	},
	)

	ast.Equal(expected, actual)




	a := []Interval{
		{
			Start: 31,
			End:   0,
		},
		{
			Start: 1,
			End:   2,
		},
		{
			Start: 3,
			End:   4,
		},
		{
			Start: 0,
			End:   4,
		},

	}
	QuickSort(a,0, len(a)-1)
	print(a[0].Start,a[0].End)
	print(a[1].Start,a[1].End)
	print(a[2].Start,a[2].End)




}
