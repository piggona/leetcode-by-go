package maxValue

import "testing"

func TestMaxValue(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	res := maxValue(grid)
	t.Errorf("%v", res)
}
