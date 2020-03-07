package basicCalculator

import "testing"

func TestBasic(t *testing.T) {
	input := "1-1+1"
	res := calculate(input)
	t.Errorf("%d", res)
}
