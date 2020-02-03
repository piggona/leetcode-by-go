package partition

import "testing"

func TestPartition(t *testing.T) {
	input := []int{1, 2, 5}
	res := canPartition(input)
	if res {
		t.Errorf("Error answer,%v", res)
	}
}
