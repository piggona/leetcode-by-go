package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	dp := make([]int, len(s))
	if s[0] == '(' && s[1] == ')' {
		dp[1] = 2
	}
	for i := 2; i < len(s); i++ {
		if s[i] == ')' && s[i-1] == '(' {
			dp[i] = dp[i-2] + 2
		} else if s[i] == ')' {
			if i-dp[i-1]-1 > 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
			} else if i-dp[i-1]-1 == 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 2
			}
		}
	}
	result := 0
	for _, num := range dp {
		if num > result {
			result = num
		}
	}
	return result
}

func main() {
	// input := bufio.NewScanner(os.Stdin)
	s := ")()())"
	// for input.Scan() {
	// 	s = input.Text()
	// }
	fmt.Println(longestValidParentheses(s))
}
