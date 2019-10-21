package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	line := len(obstacleGrid)
	col := len(obstacleGrid[0])
	bitMap := make([][]int, line)
	for i := 0; i < line; i++ {
		bitMap[i] = make([]int, col)
	}

	for i := 0; i < col && obstacleGrid[0][i] != 1; i++ {
		bitMap[0][i] = 1
	}

	for j := 0; j < line && obstacleGrid[j][0] != 1; j++ {
		bitMap[j][0] = 1
	}

	for i := 1; i < line; i++ {
		for j := 1; j < col; j++ {
			val := obstacleGrid[i][j]
			if val != 1 {
				bitMap[i][j] = bitMap[i-1][j] + bitMap[i][j-1]
			}
		}
	}
	return bitMap[line-1][col-1]
}

func main() {
	obstacleGrid := [][]int{
		{1, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
