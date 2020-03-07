package getleast

import "testing"

func TestGetLeast(t *testing.T) {
	input := []int{0, 0, 0, 2, 0, 5}
	res := getLeastNumbers(input, 0)
	t.Errorf("%v", res)
}
