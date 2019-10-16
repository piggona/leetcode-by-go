package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	ret := make([][]int, 0)
	for i, v := range intervals {
		if v[1] < newInterval[0] {
			ret = append(ret, v)
			continue
		}

		if v[0] > newInterval[1] {
			ret = append(ret, newInterval)
			ret = append(ret, intervals[i:]...)
			return ret
		}

		newInterval[0] = min(newInterval[0], v[0])
		newInterval[1] = max(newInterval[1], v[1])
	}
	return append(ret, newInterval)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	input := [][]int{
		{1, 5},
	}
	newInterval := []int{2, 3}
	// input := [][]int{
	// 	{1, 2},
	// 	{3, 5},
	// 	{6, 7},
	// 	{8, 10},
	// 	{12, 16},
	// }
	// newInterval := []int{2, 5}

	fmt.Println(insert(input, newInterval))
}
