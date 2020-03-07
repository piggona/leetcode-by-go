package straight

import "testing"

func TestStraight(t *testing.T) {
	input := []int{4, 7, 5, 9, 2}
	res := isStraight(input)
	t.Errorf("%v", res)
}
