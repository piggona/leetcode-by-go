package main

import "fmt"

func combine(n int, k int) [][]int {
	result := [][]int{}
	recur(1, n+1, k, &result, []int{})
	return result
}

func recur(pos, n, k int, result *[][]int, sel []int) {
	if len(sel) == k {
		*result = append(*result, sel)
		return
	}
	end := n - k + len(sel)
	for i := pos; i <= end; i++ {
		temp := make([]int, len(sel))
		copy(temp, sel)
		temp = append(temp, i)
		recur(i+1, n, k, result, temp)
	}
}

func main() {
	n := 0
	k := 2
	fmt.Println(combine(n, k))
}
