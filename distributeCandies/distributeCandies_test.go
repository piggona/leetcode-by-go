package disCan

import "testing"

func TestDisCan(t *testing.T) {
	res := distributeCandies(4, 3)
	t.Errorf("%v", res)
}
