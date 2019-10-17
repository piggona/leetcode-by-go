package main

import "fmt"

func lengthOfLastWord(s string) int {
	result := 0
	pos := len(s) - 1
	for pos >= 0 && s[pos] == ' ' {
		pos--
	}
	for pos >= 0 && s[pos] != ' ' {
		result++
		pos--
	}
	return result
}

func main() {
	input := ""
	fmt.Println(lengthOfLastWord(input))
}
