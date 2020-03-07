package findTarget

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	dp := make([]map[int]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make(map[int]int)
	}
	dp[0][0] = 1
	for i := 0; i < len(nums); i++ {
		tmap := dp[i]
		for key, value := range tmap {
			dp[i+1][key+nums[i]] += value
			dp[i+1][key-nums[i]] += value
		}
	}
	return dp[n][S]
}
