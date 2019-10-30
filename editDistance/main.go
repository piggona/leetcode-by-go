package main

import (
	"fmt"
	"math"
)

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, (len(word1) + 1))
	for i := 0; i < len(word1); i++ {
		dp[i] = make([]int, (len(word2) + 1))
	}
	for i := 1; i < len(word1); i++ {
		dp[i][0] = i
	}
	for i := 1; i < len(word2); i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(n1, n2, n3 int) int {
	min := math.MaxInt32
	if n1 < min {
		min = n1
	}
	if n2 < min {
		min = n2
	}
	if n3 < min {
		min = n3
	}
	return min
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}
