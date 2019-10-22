package main

import "fmt"

func minPathSum(grid [][]int) int {
	line := len(grid)
	col := len(grid[0])
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < line; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < line; i++ {
		for j := 1; j < col; j++ {
			temp := min(dp[i][j-1], dp[i-1][j])
			dp[i][j] = temp + grid[i][j]
		}
	}
	return dp[line-1][col-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	input := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(input))
}
