package subarray

import "testing"

func TestSubarray(t *testing.T) {
	input := []int{1, 1, 2, 1, -1}
	k := 3
	res := subarraySum(input, k)
	if res != 3 {
		t.Errorf("Not correct: %d", res)
	}
}
