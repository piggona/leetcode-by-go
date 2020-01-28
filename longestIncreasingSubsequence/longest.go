package longestIncreasingSubsequence

import "math"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	max := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = getMax(dp[i], dp[j]+1)
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func getMax(nums ...int) int {
	max := math.MinInt32
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
