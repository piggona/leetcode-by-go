package main

import "fmt"

func uniquePaths(m int, n int) int {
	bitMap := make([][]int, m)
	for i := 0; i < m; i++ {
		bitMap[i] = make([]int, n)
	}
	for i := 1; i < n; i++ {
		bitMap[0][i] = 1
	}
	for i := 0; i < m; i++ {
		bitMap[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			bitMap[i][j] = bitMap[i-1][j] + bitMap[i][j-1]
		}
	}
	return bitMap[m-1][n-1]
}

func main() {
	m := 3
	n := 2
	fmt.Println(uniquePaths(m, n))
}
