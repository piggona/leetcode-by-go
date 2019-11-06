package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	// lt := len(t)
	// subMap := make(map[byte]int)
	// right := 0
	// left := 0
	// cleft := 0
	// count := 0
	// for i := 0; i < lt; i++ {
	// 	subMap[t[i]] = -1
	// }
	// for i := 0; i < len(s); i++ {
	// 	if subMap[s[i]] == -1 {
	// 		subMap[s[i]] = i
	// 		count++
	// 		if count == 1 {
	// 			left = i
	// 		}
	// 		if count == lt {
	// 			right = i
	// 		}
	// 	} else if subMap[s[i]] != 0 {
	// 		if count < lt {
	// 			subMap[s[i]] = i
	// 		} else {
	// 			subMap[s[i]] = i

	// 		}
	// 	}
	// }
	return ""
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
