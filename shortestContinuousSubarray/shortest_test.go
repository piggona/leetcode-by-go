package shortestsubarray

import "testing"

func TestShortestArray(t *testing.T) {
	input := []int{2, 6, 4, 8, 10, 9, 15}
	res := findUnsortedSubarray(input)
	if res != 5 {
		t.Errorf("Bad result: %d", res)
	}
}
