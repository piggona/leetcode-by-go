package main

import "fmt"

func groupAnagrams(words []string) [][]string {
	cache := make(map[[26]byte]int)
	result := make([][]string, 0)
	for i := range words {
		bitMap := [26]byte{}
		for _, bit := range words[i] {
			bitMap[bit-'a']++
		}
		if idx, ok := cache[bitMap]; ok {
			result[idx] = append(result[idx], words[i])
		} else {
			result = append(result, []string{words[i]})
			cache[bitMap] = len(result) - 1
		}
	}
	return result
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(input))
}
