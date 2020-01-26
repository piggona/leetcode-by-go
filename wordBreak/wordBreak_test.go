package wordBreak

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	s := "leetcode"
	wordDict := []string{
		"leet",
		"code",
	}
	fmt.Println(wordBreakDP(s, wordDict))
}
