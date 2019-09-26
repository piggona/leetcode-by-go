package main

import "fmt"

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		}
		if s[i] == '}' || s[i] == ']' || s[i] == ')' {
			if len(stack) == 0 {
				return false
			}
			temp := stack[len(stack)-1]
			if s[i]-temp > 2 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	s := "{[dasf]}"
	fmt.Println(isValid(s))
}
