package main

import (
	"fmt"
	"math"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{[]int{}}
	numMap := make(map[string]bool)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		reslen := len(res)
		for j := 0; j < reslen; j++ {
			temp := make([]int, len(res[j]))
			copy(temp, res[j])
			new := append(temp, nums[i])
			str := fmt.Sprintf("%v", new)
			if _, ok := numMap[str]; !ok {
				res = append(res, new)
				numMap[str] = true
			}
		}
	}
	return res
}

func subsetsWithDupFine(nums []int) [][]int {
	res := [][]int{[]int{}}
	last := math.MinInt32
	size := 1
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		reslen := len(res)
		if nums[i] != last {
			last = nums[i]
			size = len(res)
		}
		newSize := len(res)
		for j := newSize - size; j < reslen; j++ {
			temp := make([]int, len(res[j]))
			copy(temp, res[j])
			new := append(temp, nums[i])
			res = append(res, new)
		}
	}
	return res
}

func main() {
	input := []int{1, 2, 2, 1}
	fmt.Println(subsetsWithDupFine(input))
}
