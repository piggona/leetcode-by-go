package main

import "fmt"

func longestValidParentheses(s string) int {
	result := 0
	pos := 0
	stack := []int{}
	for pos < len(s) && s[pos] != '(' {
		pos++
	}
	if pos == len(s) {
		return 0
	}
	s = s[pos:]
	pos = 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if len(stack) != 0 && s[i] == ')' {
			stack = stack[:len(stack)-1]
		} else {
			if len(s) > i+1 {
				result = longestValidParentheses(s[i+1:])
			}
			pos = i
			break
		}
	}
	if len(stack) == 0 {
		if result < pos {
			result = pos
		}
		if pos == 0 {
			result = len(s)
		}
	} else {
		length := stack[0]
		if result < length {
			result = length
		}
		for j := 0; j < len(stack)-1; j++ {
			length := stack[j+1] - stack[j] - 1
			if result < length {
				result = length
			}
		}
		length = len(s) - stack[len(stack)-1] - 1
		if result < length {
			result = length
		}
	}
	return result
}

func main() {
	s := ")()(("
	fmt.Println(longestValidParentheses(s))
}
