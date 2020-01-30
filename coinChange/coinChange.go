package coinChange

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		min := math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] >= 0 && min > dp[i-coin] {
				min = dp[i-coin]
			}
		}
		if min == math.MaxInt32 {
			dp[i] = -1
		} else {
			dp[i] = min + 1
		}
	}

	return dp[amount]
}
