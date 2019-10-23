package main

import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return false
	}
	// mstr := []string{"?-", "#n", "?(.*n)", "?(e*n)"}
	m := 0
	i := 0
	// isNum := true
	hasNum := false
	hasDot := false
	hasE := false
	hasSlash := false

	for i < len(s) {
		switch m {
		case 0:
			if s[i] == '-' || s[i] == '+' {
				i++
			}
			m++
		case 1:
			if (s[i] > '9' || s[i] < '0') && !hasNum {
				return false
			}
			if !hasNum {
				hasNum = true
				i++
			} else {
				m++
			}
		case 2:
			if s[i] == '.' && (i+1 == len(s) || (s[i+1] > '9' || s[i+1] < '0')) {
				return false
			}
			if s[i] != '.' && (s[i] > '9' || s[i] < '0') {
				m++
			} else if hasDot && (s[i] <= '9' && s[i] >= '0') {
				i++
			} else {
				hasDot = true
				i++
			}
		case 3:
			// if s[i] != 'e' && (s[i] > '9' || s[i] < '0') {
			// 	return false
			// }
			if s[i] == 'e' && i+1 == len(s) {
				return false
			}
			if s[i] == 'e' {
				hasE = true
				i++
			} else if hasE && !hasSlash && (s[i] == '-' || s[i] == '+') {
				hasSlash = true
				i++

			} else if hasE && (s[i] <= '9' && s[i] >= '0') {
				i++
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	s := "95a54e53"
	fmt.Println(isNumber(s))
}
