package reversePairs

import "testing"

func TestPairs(t *testing.T) {
	input := []int{7, 5, 6, 4}
	res := reversePairs(input)
	t.Errorf("%v", res)
}
