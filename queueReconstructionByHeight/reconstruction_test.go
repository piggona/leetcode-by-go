package reconstruction

import "testing"

func TestReconstruction(t *testing.T) {
	input := [][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}
	res := ReconstructQueue(input)
	t.Errorf("res:%v", res)
}
