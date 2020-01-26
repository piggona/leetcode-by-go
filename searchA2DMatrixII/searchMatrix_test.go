package searchA2DMatrixII

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	input := [][]int{
		{-5},
		// {1, 4, 7, 11, 15},
		// {2, 5, 8, 12, 19},
		// {3, 6, 9, 16, 22},
		// {10, 13, 14, 17, 24},
		// {18, 21, 23, 26, 30},
	}
	target := -10
	res := searchMatrix(input, target)
	if res {
		t.Errorf("Bad result: %v", res)
	}
}
