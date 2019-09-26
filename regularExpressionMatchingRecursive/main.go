package main

import "fmt"

func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}
	if s == "" && p == "" {
		return true
	}
	if s != "" && p == "" {
		return false
	}
	matchList := getSplit(p)
	return getMatch(s, matchList)
}

func getMatch(s string, p []string) bool {
	str := p[0]
	i := 0
	if str[len(str)-1] == '*' {
		if s == "" {
			if len(p) == 1 {
				return true
			}
			return getMatch(s, p[1:])
		}
		if len(p) == 1 && len(s) == 1 {
			if s[0] == str[0] || str[0] == '.' {
				return true
			} else {
				return false
			}
		}
		if len(p) == 1 && len(s) != 1 {
			if s[0] == str[0] || str[0] == '.' {
				return getMatch(s[1:], p)
			} else {
				return false
			}
		}
		if len(p) != 1 && len(s) == 1 {
			if s[0] == str[0] || str[0] == '.' {
				return getMatch(s, p[1:]) || getMatch("", p)
			}
			return getMatch(s, p[1:])
		}
		if s[0] == str[0] || str[0] == '.' {
			return getMatch(s, p[1:]) || getMatch(s[1:], p)
		} else {
			return getMatch(s, p[1:])
		}
	} else {
		if s == "" {
			return false
		}
		for i = 0; i < len(str); i++ {
			if i >= len(s) {
				return false
			}
			if s[i] != str[i] && str[i] != '.' {
				return false
			}
		}
		if i == len(s) && len(p) == 1 {
			return true
		}
		if i == len(s) && len(p) != 1 {
			return getMatch("", p[1:])
		}
		if i != len(s) && len(p) == 1 {
			return false
		}
		return getMatch(s[i:], p[1:])
	}
}

func getSplit(p string) []string {
	result := []string{}
	for i := 1; i < len(p); i++ {
		if p[i] == '*' {
			if i-2 >= 0 {
				result = append(result, p[:i-1])
			}
			if i+1 < len(p) {
				result = append(result, p[i-1:i+1])
				p = p[i+1:]
			} else {
				result = append(result, p[i-1:])
				return result
			}
			i = 0
		}
	}
	result = append(result, p)
	return result
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
