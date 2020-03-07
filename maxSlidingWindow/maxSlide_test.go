package maxSlide

import "testing"

func TestSlide(t *testing.T) {
	nums := []int{1}
	k := 1
	res := maxSlidingWindow(nums, k)
	t.Errorf("%v", res)
}
