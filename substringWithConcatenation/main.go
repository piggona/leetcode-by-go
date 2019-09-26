package main

import (
	"fmt"
	"strings"
)

func compareSub(sub string, words []string) bool {
	twords := make([]string, len(words))
	copy(twords, words)
	if len(words) == 0 {
		return true
	}
	if len(sub) == 0 {
		return false
	}
	for id, word := range words {
		start := strings.Index(sub, word)
		if start == 0 {
			twords[0], twords[id] = twords[id], twords[0]
			if len(sub) == start+len(word) {
				if len(words) == 1 {
					return true
				}
			}
			if compareSub(sub[start+len(word):], twords[1:]) {
				return true
			}
		}
	}
	return false
}

func findSubstring(s string, words []string) []int {
	result := []int{}
	w := make(map[string]int)
	for id, word := range words {
		w[word] = id
	}
	for word, id := range w {
		temp := s
		if strings.Index(s, word) == -1 {
			return []int{}
		}
		start := strings.Index(s, word)
		twords := make([]string, len(words))
		copy(twords, words)
		twords[0], twords[id] = twords[id], twords[0]
		for start != -1 {
			if len(temp) == start+len(word) && len(twords) == 1 {
				result = append(result, start)
			} else if compareSub(temp[start+len(word):], twords[1:]) {
				result = append(result, start)
			}

			t := []byte(temp)
			t[start] = '#'
			temp = string(t)
			start = strings.Index(temp, word)
		}
	}
	return result
}

func main() {
	s := "aaaaaaaa"
	words := []string{"aa", "aa", "aa"}
	// s := "wordgoodgoodgoodbestword"
	// words := []string{"word", "good", "best", "good"}
	// s := "barfoofoobarthefoobarman"
	// words := []string{"bar", "foo", "the"}
	// s := ""
	// words := []string{"rich"}
	fmt.Println(findSubstring(s, words))
}
