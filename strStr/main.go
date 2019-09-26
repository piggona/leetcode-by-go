package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	partMatch := getPartMatch(needle)
	hpos := 0
	npos := 0
	for hpos < len(haystack) && npos < len(needle) {
		if haystack[hpos] == needle[npos] {
			hpos++
			npos++
		} else {
			if npos == 0 {
				hpos++
				npos = 0
			} else {
				hpos = hpos - partMatch[npos-1]
				npos = 0
			}
		}
	}
	if npos == len(needle) {
		return hpos - npos
	} else {
		return -1
	}

}

func getPartMatch(needle string) []int {
	result := make([]int, len(needle))
	result[0] = 0
	i := 0
	j := 1
	for j < len(needle) {
		if needle[i] == needle[j] {
			result[j] = i + 1
			i++
			j++
		} else {
			if i != 0 {
				i = result[i-1]
			} else {
				j++
			}
		}
	}
	return result
}

func failTable(p string) []int {
	l := len(p)
	t := make([]int, l)
	t[0] = 0
	i, j := 0, 1
	for j < l {
		if p[i] == p[j] {
			t[j] = i + 1
			i++
			j++
		} else {
			if i != 0 {
				i = t[i-1]
			} else {
				t[j] = i
				j++
			}
		}
	}
	return t
}

func main() {
	// haystack := "mississipi"
	needle := "aabaaac"
	// fmt.Println(strStr(haystack, needle))
	fmt.Println(getPartMatch(needle))
	fmt.Println(failTable(needle))
}
