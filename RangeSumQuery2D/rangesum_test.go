package RangeSumQuery2D

import (
	"fmt"
	"testing"
)

func TestRangeSum(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	st := Constructor(matrix)
	res := st.SumRegion(2, 1, 4, 3)
	fmt.Println(res)
}
