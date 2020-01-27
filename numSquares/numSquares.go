package numSquares

import (
	"math"
)

func numSquaresDI(n int) int {
	for n%4 == 0 {
		n = n / 4
	}
	if n%8 == 7 {
		return 4
	}
	for i := 0; i*i <= 0; i++ {
		b := math.Sqrt(float64(n - i*i))
		if float64(int(b)) == b && i*i+int(b*b) == n {
			if i >= 0 && b >= 0 {
				return 2
			}
			return 1
		}
	}
	return 3
}

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	for n%4 == 0 {
		n = n / 4
	}
	if n%8 == 7 {
		return 4
	}
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = getMin(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
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
