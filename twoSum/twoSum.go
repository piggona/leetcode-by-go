package twoSum

import "math"

func twoSum(n int) []float64 {
	const max = 6
	var (
		total  = math.Pow(float64(max), float64(n))
		result []float64
		dp     [][]int
	)
	dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n*max+1)
	}
	for i := 1; i < len(dp[0]); i++ {
		dp[1][i] = 1
	}
	for i := 2; i < n+1; i++ {
		for j := i; j < n*max+1; j++ {
			for k := 1; k < 7; k++ {
				if j-k > 0 {
					dp[i][j] += dp[i-1][j-k]
				}
			}
		}
	}
	for _, val := range dp[n] {
		if val != 0 {
			result = append(result, float64(val)/total)
		}
	}
	return result

}
