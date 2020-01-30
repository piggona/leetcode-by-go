package coinChange

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	res := coinChange(coins, amount)
	if res != 3 {
		t.Errorf("Not Correct:%v", res)
	}
}
