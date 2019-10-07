package main

import (
	"fmt"
	"sort"
)

type sortByte []byte

func (b sortByte) Len() int {
	return len(b)
}

func (b sortByte) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b sortByte) Swap(i, j int) {
	b[j], b[i] = b[i], b[j]
}

func groupAnagrams(str []string) [][]string {
	if len(str) == 0 {
		return [][]string{}
	}
	strMap := make(map[string][]int)
	for id, s := range str {
		t := sortByte([]byte(s))
		sort.Sort(t)
		st := string(t)
		if strMap[st] == nil {
			temp := []int{id}
			strMap[st] = temp
		} else {
			temp := strMap[st]
			temp = append(temp, id)
			strMap[st] = temp
		}
	}
	result := [][]string{}
	for _, val := range strMap {
		temp := make([]string, len(val))
		for id, m := range val {
			temp[id] = str[m]
		}
		result = append(result, temp)
	}
	return result

}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(input))
}
