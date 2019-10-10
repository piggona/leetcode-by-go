package main

import "fmt"

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	helper(nums, []int{}, &result)
	return result
}

func helper(candidate []int, temp []int, result *[][]int) {
	if len(candidate) == 0 {
		*result = append(*result, temp)
		return
	}
	for id, num := range candidate {
		ct := append([]int{}, candidate...)
		tt := append([]int{}, temp...)
		helper(append(ct[:id], ct[id+1:]...), append(tt, num), result)
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
