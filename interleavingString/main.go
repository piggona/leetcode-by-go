package main

import "fmt"

func isInterleaveBad(s1 string, s2 string, s3 string) bool {
	if len(s1) != len(s2)+len(s3) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	result := false
	for i := 0; i < 2; i++ {
		if i == 0 && len(s2) != 0 && s1[0] == s2[0] {
			result = result || isInterleave(s1[1:], s2[1:], s3)
		} else if i == 1 && len(s3) != 0 && s1[0] == s3[0] {
			result = result || isInterleave(s1[1:], s2, s3[1:])
		}
	}
	return result
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 1; i < len(s1)+1; i++ {
		if dp[i-1][0] == true && s1[i-1] == s3[i-1] {
			dp[i][0] = true
		}
	}
	for i := 1; i < len(s2)+1; i++ {
		if dp[0][i-1] == true && s2[i-1] == s3[i-1] {
			dp[0][i] = true
		}
	}
	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	fmt.Println(dp)
	return dp[len(s1)][len(s2)]
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	fmt.Println(isInterleave(s1, s2, s3))
}
