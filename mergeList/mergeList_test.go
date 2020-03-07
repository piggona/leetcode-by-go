package mergeList

import "testing"

func TestMergeList(t *testing.T) {
	A := []int{1, 2, 3, 0, 0, 0}
	n := 3
	B := []int{2, 5, 6}
	m := 3
	merge(A, n, B, m)
	t.Errorf("%v", A)
}
