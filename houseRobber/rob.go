package houseRobber

import "math"

func robRecur(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) <= 2 {
		return getMax(nums...)
	}
	if len(nums) == 3 {
		sum1 := nums[0] + nums[2]
		return getMax(sum1, nums[1])
	}
	sum1 := nums[0] + rob(nums[2:])
	sum2 := nums[1] + rob(nums[3:])
	return getMax(sum1, sum2)
}

func getMax(max ...int) int {
	res := math.MinInt32
	for _, value := range max {
		if value > res {
			res = value
		}
	}
	return res
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) <= 2 {
		return getMax(nums...)
	}
	if len(nums) == 3 {
		sum1 := nums[0] + nums[2]
		return getMax(sum1, nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = nums[2] + nums[0]
	for i := 3; i < len(nums); i++ {
		dp[i] = nums[i] + getMax(dp[i-2], dp[i-3])
	}
	return dp[len(nums)-1]
}
