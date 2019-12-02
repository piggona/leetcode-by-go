package main

import (
	"fmt"
	"strconv"
	"strings"
)

func numDecodings(s string) int {
	result := 0
	if strings.Index(s, "0") != -1 {
		return 0
	}
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	if len(s) == 2 {
		num, _ := strconv.Atoi(s)
		if num <= 26 {
			return 2
		}
		return 1
	}
	for i := 1; i <= 2; i++ {
		temp, _ := strconv.Atoi(s[:i])
		if temp <= 26 {
			result += numDecodings(s[i:])
		}
	}
	return result
}

func numDecodingsDp(s string) int {
	if s[0] == '0' {
		return 0
	}

	n := len(s)
	dp := make([]int, n+1)

	dp[0], dp[1] = 1, 1
	lastOne, lastTwo := int(s[0]-'0'), 0
	for i := 2; i <= n; i++ {
		last := int(s[i-1] - '0')
		lastOne, lastTwo = last, lastOne*10+last
		if 1 <= lastOne {
			dp[i] = dp[i-1]
		}
		if 10 <= lastTwo && lastTwo <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func main() {
	input := "262626"
	fmt.Println(numDecodings(input))
}
