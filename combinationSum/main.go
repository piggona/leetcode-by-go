package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	if len(candidates) == 0 {
		return result
	}
	sort.Ints(candidates)
	temp := []int{}
	result = combinate(result, temp, candidates, target)
	return result
}

func combinate(result [][]int, temp []int, candidates []int, target int) [][]int {
	for i := 0; i < len(candidates); i++ {
		count := 1
		for count*candidates[i] < target {
			temp = append(temp, candidates[i])
			count++
		}
		if count*candidates[i] == target {
			temp = append(temp, candidates[i])
			t := make([]int, len(temp))
			copy(t, temp)
			result = append(result, t)
			count = count - 1
		} else {
			count = count - 2
		}
		for count > 0 {
			temp = temp[:len(temp)-1]
			result = combinate(result, temp, candidates[i+1:], target-count*candidates[i])
			count -= 1
		}
		if count == 0 {
			temp = temp[:len(temp)-1]
		}
	}
	return result
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
