package main

import (
	"fmt"
	"strings"
)

var M = map[byte]string{
	'1': "",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
	'0': " ",
}

func letterCombinations(digits string) []string {
	ans := []string{}
	if len(digits) == 0 {
		return ans
	}
	combiner(digits, 0, "", &ans)
	return ans
}

func combiner(digits string, pos int, s string, ans *[]string) {
	if pos == len(digits) {
		*ans = append(*ans, s)
		return
	}
	for _, c := range M[digits[pos]] {
		combiner(digits, pos+1, s+string(c), ans)
	}
	return
}

func combine(a []string, b []string) []string {
	result := []string{}
	var str strings.Builder
	for _, i := range a {
		for _, j := range b {
			str.WriteString(i)
			str.WriteString(j)
			result = append(result, str.String())
		}
	}
	return result
}

func combineByte(a *[]string, b byte) *[]string {
	var str strings.Builder
	for _, i := range *a {
		str.WriteString(i)
		str.WriteByte(b)
		*a = append(*a, str.String())
	}
	return a
}

func main() {
	input := "23"
	fmt.Println(letterCombinations(input))
}
