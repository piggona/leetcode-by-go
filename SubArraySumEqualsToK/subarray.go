package subarray

func subarraySumBad(nums []int, k int) int {
	dp := make([][]int, len(nums))
	result := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}
	for i := 0; i < len(nums); i++ {
		dp[i][i] = nums[i]
		if dp[i][i] == k {
			result++
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			dp[i][j] = dp[i][j-1] + nums[j]
			if dp[i][j] == k {
				result++
			}
		}
	}
	return result
}

func subarraySum(nums []int, k int) int {
	sums := make([]int, len(nums))
	result := 0
	copy(sums, nums)
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if sums[i] == k {
			result++
		}
		for j := i - 1; j >= 0; j-- {
			if sums[i]-sums[j] == k {
				result++
			}
		}
	}
	return result
}
