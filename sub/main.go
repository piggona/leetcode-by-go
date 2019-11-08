package main

import (
	"fmt"
	"sort"
)

func subsets(num []int) [][]int {
	result := [][]int{[]int{}}
	length := len(num)
	sort.Ints(num)
	for i := 0; i < length; i++ {
		reslen := len(result)
		for j := 0; j < reslen; j++ {
			temp := make([]int, len(result[j]))
			copy(temp, result[j])
			temp = append(temp, num[i])
			result = append(result, temp)
		}
	}
	return result
}

func recur(nums []int, temp []int, start int, need int, result *[][]int) {
	if need == 0 {
		*result = append(*result, temp)
		return
	}
	for i := start; i <= len(nums)-need; i++ {
		t := append(temp, nums[i])
		recur(nums, t, i+1, need-1, result)
	}
}

func main() {
	nums := []int{9, 0, 3, 5, 7}
	fmt.Println(subsets(nums))
}
