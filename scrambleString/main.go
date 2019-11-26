package main

import (
	"fmt"
	"strings"
)

func isScrambleNotAvailable(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		if len(s1) == 0 && len(s2) == 0 {
			return true
		} else {
			return false
		}
	}
	if len(s1) == 1 {
		return s1 == s2
	}
	if len(s1) == 2 {
		return sameDistribute2(s1, s2)
	}
	div := string(s2[0])
	var temp = s1
	former := 0
	res := false
	pos := strings.Index(temp, div)
	for pos != -1 {
		left := s1[0 : former+pos]
		right := s1[former+pos+1:]
		left2 := s2[1 : len(left)+1]
		right2 := s2[len(left)+1:]
		left3 := s2[1 : len(right)+1]
		right3 := s2[len(right)+1:]
		res = res || (isScramble(left, left2) && isScramble(right, right2)) || (isScramble(left, right3) && isScramble(right, left3))
		if res == true {
			return true
		}
		former = pos
		temp = temp[pos+1:]
		pos = strings.Index(temp, div)
	}
	return false
}

func sameDistribute2(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if s1[0] == s2[1] && s2[0] == s1[1] {
		return true
	}
	return false
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	s1Distribute := make(map[byte]int)
	s2Distribute := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		s1Distribute[s1[i]]++
		s2Distribute[s2[i]]++
	}

	if !sameDistribute(s1Distribute, s2Distribute) {
		return false
	}

	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i]) {
			return true
		}
	}
	return false
}

func sameDistribute(s1, s2 map[byte]int) bool {
	for key, value := range s1 {
		if s2[key] != value {
			return false
		}
	}
	return true
}

func main() {
	s1 := "rgeat"
	s2 := "great"
	fmt.Println(isScramble(s1, s2))
}
