package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if strs == nil {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if !strings.HasPrefix(strs[i], prefix) {
			str := strs[i]
			j := 0
			for j = 0; j < len(prefix) && j < len(str); j++ {
				if prefix[j] != str[j] {
					break
				}
			}
			if j == 0 {
				return ""
			}
			prefix = prefix[:j]
		}
	}
	return prefix
}

func main() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}
