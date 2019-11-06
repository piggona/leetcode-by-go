package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	need := make(map[byte]int, len(t))
	have := make(map[byte]int, len(t))
	left := 0
	res := ""
	cnt := 0
	minLen := math.MaxInt32

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	for i := 0; i < len(s); i++ {
		if have[s[i]] < need[s[i]] {
			cnt++
		}
		have[s[i]]++
		for left <= i && have[s[left]] > need[s[left]] {
			have[s[left]]--
			left++
		}

		length := i - left + 1
		if cnt == len(t) && minLen > length {
			minLen = length
			res = s[left : i+1]
		}
	}
	return res
}

// func minWindow(s string, t string) string {
// 	subMap := make(map[byte]int)
// 	tlen := len(t)
// 	res := 0
// 	cnt := 0
// 	minLen := math.MaxInt32
// 	left := 0
// 	right := 0
// 	for i := 0; i < len(t); i ++ {
// 		subMap[t[i]]++
// 	}
// 	pos := 0
// 	for pos = 0 ; pos < len(s); pos ++ {
// 		if cnt == tlen {
// 			right = pos
// 			res = s[left:right+1]
// 			minLen = right - left
// 			if subMap[s[left]] >= 0 {
// 				cnt --
// 				subMap[s[left]] ++
// 			}
// 			left ++
// 		}else {
// 			if map[s[pos]] < 1
// 		}
// 	}

// }

func getBorder(sub *map[byte]int) (int, int) {
	min := math.MaxInt32
	max := math.MinInt32
	for _, i := range *sub {
		if i < min {
			min = i
		} else if i > max {
			max = i
		}
	}
	return min, max
}

func main() {
	S := "ADOBECODEBANC"
	T := "ABC"
	fmt.Println(minWindow(S, T))
}
