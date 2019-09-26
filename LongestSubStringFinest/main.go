package main

import "fmt"

func lengthOfLongestSubString(s string) int {
	array := [256]int{}
	for i := 0; i < len(array); i++ {
		array[i] = -1
	}
	pre := -1
	cur := 0
	length := 0
	for i := 0; i < len(s); i++ {
		pre = max(pre, array[s[i]])
		cur = i - pre
		length = max(length, cur)
		array[s[i]] = i
	}
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	s := "abcabcbb"
	result := lengthOfLongestSubString(s)
	fmt.Println(result)
}
