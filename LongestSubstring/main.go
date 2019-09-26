package main

import "fmt"

func lengthOfLongestSubString(s string) int {
	str := []byte(s)
	result := 0
	if len(str) < 2 {
		return len(str)
	}
	for i := 0; i < len(str); i++ {
		count := 1
		hashSet := make(map[byte]int)
		hashSet[str[i]] = i + 1
		for j := i + 1; j < len(str); j++ {
			if hashSet[str[j]] != 0 {
				break
			}
			count++
			hashSet[str[j]] = j + 1
		}
		if count > result {
			result = count
		}
	}
	return result
}

func main() {
	s := "abcabcbb"
	result := lengthOfLongestSubString(s)
	fmt.Println(result)
}
