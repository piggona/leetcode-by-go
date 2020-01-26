package productOfArrayExpectSelf

import "testing"

func TestProduct(t *testing.T) {
	input := []int{1, 2, 3, 4}
	res := productExceptSelf(input)
	correct := []int{24, 12, 8, 6}
	for i := 0; i < len(res); i++ {
		if res[i] != correct[i] {
			t.Errorf("position %d not correct.The wrong answer is %v", i, res)
		}
	}
}
