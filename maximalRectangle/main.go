package main

import (
	"fmt"
	"math"
)

func maximalRectangle(matrix [][]byte) int {
	line := len(matrix) - 1
	col := len(matrix[0]) - 1
	return Rectangle(matrix, 0, 0, line, col)
}

func Rectangle(matrix [][]byte, lpos, cpos, line, col int) int {
	for i := lpos; i <= line; i++ {
		for j := cpos; j <= col; j++ {
			if matrix[i][j] == '0' {
				var down, left, right int
				if i == line {
					down = 0
				} else {
					down = Rectangle(matrix, i+1, cpos, line, col)
				}
				if j == cpos {
					left = 0
				} else {
					left = Rectangle(matrix, lpos, cpos, line, j-1)
				}
				if j == col {
					right = 0
				} else {
					right = Rectangle(matrix, lpos, j+1, line, col)
				}
				return getMax(down, left, right)
			}
		}
	}
	return (line - lpos + 1) * (col - cpos + 1)
}

func getMax(nums ...int) int {
	max := math.MinInt32
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	input := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	fmt.Println(maximalRectangle(input))
}
