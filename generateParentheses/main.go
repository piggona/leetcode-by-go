package main

import "fmt"

func generateParenthesis(n int) []string {
	stack := []byte{}
	result := &[]string{}
	if n == 0 {
		return *result
	}
	identify := make(map[string]bool)
	parenthesis(n, "", stack, result, identify)
	return *result
}

func parenthesis(n int, str string, stack []byte, result *[]string, identify map[string]bool) {
	if n == 0 {
		temp := []byte{}
		for i := 0; i < len(stack); i++ {
			temp = append(temp, ')')
		}
		str = str + string(temp)
		if !identify[str] {
			*result = append(*result, str)
			identify[str] = true
		}
		return
	}
	if len(stack) == 0 {
		parenthesis(n-1, str+"()", stack, result, identify)
	}
	for i := 0; i <= len(stack); i++ {
		stk := stack
		temp := []byte{}
		for j := 0; j < i; j++ {
			temp = append(temp, ')')
		}
		stk = stk[:len(stk)-i]
		parenthesis(n-1, str+"()"+string(temp), stk, result, identify)
	}
	stack = append(stack, '(')
	parenthesis(n-1, str+"(", stack, result, identify)
}

func main() {
	n := 3
	result := []string{}
	result = generateParenthesis(n)
	fmt.Println(result)
}
