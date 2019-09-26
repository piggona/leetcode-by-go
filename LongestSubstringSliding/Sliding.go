package main

import "fmt"

func lengthOfLongestSubString(s string) int {
	result := 0
	hashSet := make(map[byte]int)
	left, right := 0, 0
	for right < len(s) {
		temp := s[right]
		if val, ok := hashSet[temp]; !ok || (ok && val == 0) {
			hashSet[temp] = right + 1
			if right-left+1 > result {
				result = right - left + 1
			}
			right++
			continue
		}

		temp = s[left]
		left++
		if val, ok := hashSet[temp]; ok {
			hashSet[temp] = val - 1
		}
	}
	return result
}

func main() {
	s := "abcabcbb"
	result := lengthOfLongestSubString(s)
	fmt.Println(result)
}
