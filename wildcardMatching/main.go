package main

import "fmt"

func isMatch(s string, p string) bool {
	// mem[i][j] means isMatch(s[:i], p[:j])
	mem := make([][]bool, len(s)+1)
	for i := range mem {
		mem[i] = make([]bool, len(p)+1)
	}

	// init bound, mem[n][0] is false while n > 0
	mem[0][0] = true
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			mem[0][j] = mem[0][j-1]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				mem[i][j] = mem[i][j-1] || mem[i-1][j]
			} else if p[j-1] == '?' || p[j-1] == s[i-1] {
				mem[i][j] = mem[i-1][j-1]
			}
		}
	}

	return mem[len(s)][len(p)]
}

func main() {
	s := "mississippi"
	p := "mi?s*p*"
	// s := "adceb"
	// p := "*a*b"
	// s := "a"
	// p := "aa"
	// s := "ab"
	// p := "*?*?*"
	fmt.Println(isMatch(s, p))
}
