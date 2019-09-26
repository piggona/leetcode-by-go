package main

import "fmt"

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	if m > 0 && n > 0 {
		dp[1][1] = s[0] == p[0] || p[0] == '.'
	}
	for j := 2; j <= m; j++ {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}
	for i := 1; i <= n; i++ {
		for j := 2; j <= m; j++ {
			if p[j-1] != '*' {
				dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			} else {
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			}
		}
	}
	return dp[n][m]
}

func main() {
	// s := "mississippi"
	// p := "mis*is*ip*."
	// s := "aaabdzabc"
	// p := "c*a*.*bc"
	// s := "baaaabaaab"
	// p := "b.*aba*ab"
	// s := "asdfaaba"
	// p := "asdfb*.*c*a"
	// s := ""
	// p := ".*"
	// s := "aab"
	// p := "c*a*b"
	// s := "a"
	// p := "ab*"
	s := "abbabaaaaaaacaa"
	p := "a*.*b.a.*c*b*a*c*"
	// s := ""
	// p := "c*"
	fmt.Println(isMatch(s, p))
	// fmt.Println(getSplit(p))
}
