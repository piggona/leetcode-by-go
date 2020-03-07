package translateNum

import "testing"

func TestTranslate(t *testing.T) {
	input := 506
	res := translateNum(input)
	t.Errorf("%v", res)
}
