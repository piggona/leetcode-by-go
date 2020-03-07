package translateNum

import "strconv"

func translateNum(num int) int {
	var (
		nums = strconv.Itoa(num)
		dp   = make([]int, len(nums)+1)
	)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i <= len(nums)-1; i++ {
		temp, _ := strconv.Atoi(nums[i-1 : i+1])
		if nums[i-1]!= '0' && temp < 26 {
			dp[i+1] = dp[i] + dp[i-1]
		} else {
			dp[i+1] = dp[i]
		}
	}
	return dp[len(nums)]
}
