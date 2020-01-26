package maximumProductSubArray

import "math"

func maxProductBad(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	pos := 1
	neg := -1
	max := math.MinInt32
	dp := make([]int, len(nums))
	if nums[0] < 0 {
		neg = 0
		temp := maxProduct(nums[1:])
		if temp > max {
			max = temp
		}
	}
	if nums[0] == 0 {
		temp := maxProduct(nums[1:])
		if temp > 0 {
			return getMax(temp, max)
		}
		return getMax(0, max)
	}
	dp[0] = nums[0]
	for pos < len(nums) {
		if neg == -1 && nums[pos] < 0 {
			neg = pos
			if pos+1 < len(nums) {
				temp := maxProduct(nums[pos+1:])
				if temp > max {
					max = temp
				}
			}
		} else if neg != -1 && nums[pos] < 0 {
			temp := product(nums, neg, pos)
			if neg == 0 {
				dp[pos] = temp
			} else {
				dp[pos] = dp[neg-1] * temp
			}
			neg = -1
			pos++
			continue
		}
		if nums[pos] == 0 {
			temp := maxProduct(nums[pos+1:])
			if dp[pos-1] > 0 {
				if temp > dp[pos-1] {
					return getMax(temp, max)
				}
				return getMax(dp[pos-1], max)
			} else {
				if temp > 0 {
					return getMax(temp, max)
				}
				return getMax(0, max)
			}
		}
		if neg == -1 {
			dp[pos] = dp[pos-1] * nums[pos]
		} else {
			dp[pos] = dp[pos-1]
		}
		pos++
	}
	return getMax(dp[pos-1], max)
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

func getMin(nums ...int) int {
	res := math.MaxInt32
	for _, value := range nums {
		if value < res {
			res = value
		}
	}
	return res
}

func maxProduct(nums []int) int {
	res := nums[0]
	f := make([]int, len(nums))
	g := make([]int, len(nums))
	f[0] = nums[0]
	g[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		f[i] = getMax(f[i-1]*nums[i], g[i-1]*nums[i], nums[i])
		g[i] = getMin(f[i-1]*nums[i], g[i-1]*nums[i], nums[i])
		res = getMax(res, f[i])
	}
	return res
}

func product(nums []int, start int, end int) int {
	res := 1
	for i := start; i <= end; i++ {
		res = res * nums[i]
	}
	return res
}
