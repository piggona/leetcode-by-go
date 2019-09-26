package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	if len(candidates) == 0 {
		return result
	}
	temp := []int{}
	sort.Ints(candidates)
	unique := make(map[string]bool)
	return combinate(result, candidates, temp, target, unique)

}

func combinate(result [][]int, candidates []int, temp []int, target int, unique map[string]bool) [][]int {
	for i := 0; i < len(candidates); i++ {
		if candidates[i] == target {
			temp = append(temp, candidates[i])
			if unique[fmt.Sprint(temp)] == false {
				unique[fmt.Sprint(temp)] = true
				t := make([]int, len(temp))
				copy(t, temp)
				result = append(result, t)
			}
			break
		} else if target-candidates[i] >= candidates[i] {
			temp = append(temp, candidates[i])
			result = combinate(result, candidates[i+1:], temp, target-candidates[i], unique)
			temp = temp[:len(temp)-1]
		}
	}
	return result
}

func main() {
	candidates := []int{4, 4, 2, 1, 4, 2, 2, 1, 3}
	target := 6
	fmt.Println(combinationSum2(candidates, target))
}
