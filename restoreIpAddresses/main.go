package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	result := []string{}
	count := 4
	iprecur(s, -1, count, &result)
	return result
}

func iprecur(s string, pos int, count int, result *[]string) {
	if count == 1 {
		tstr := s[pos+1:]
		if available(tstr) {
			*result = append(*result, s)
			return
		}
		return
	}
	// if count == 0 && pos == len(s)-1 {
	// 	*result = append(*result, s[:pos])
	// 	return
	// }
	// if count == 0 || pos == len(s) {
	// 	return
	// }
	for i := 1; i < 4; i++ {
		tpos := pos + 1
		tpos += i
		if tpos < len(s) {
			tstr := s[pos+1 : tpos]
			if !available(tstr) {
				continue
			}
			temp := s[:tpos] + "." + s[tpos:]
			iprecur(temp, tpos, count-1, result)
		} else if tpos == len(s) {
			tstr := s[pos+1:]
			if !available(tstr) {
				continue
			}
			temp := s[:tpos] + "."
			iprecur(temp, tpos, count-1, result)
		}
	}
}

func available(tstr string) bool {
	if len(tstr) == 0 {
		return false
	}
	if num, _ := strconv.Atoi(tstr); num > 255 {
		return false
	}
	if tstr[0] == '0' && len(tstr) > 1 {
		return false
	}
	return true
}

func main() {
	input := "25525511135"
	fmt.Println(restoreIpAddresses(input))
}
