package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		check := intervals[i]
		if cur[1] >= check[0] {
			cur[1] = max(cur[1], check[1])
		} else {
			result = append(result, cur)
			cur = check
		}
	}
	result = append(result, cur)
	return result
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func main() {
	input := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(merge(input))
}
