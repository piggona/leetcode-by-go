package isNumber

import "testing"

func TestIsNumber(t *testing.T) {
	input := "+.8"
	res := isNumber(input)
	t.Errorf("%v", res)
}
