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

func largestRectangleAreaStack(height []int) int {
	height = append(height, 0)
	stack := []int{}
	i := 0
	max := 0
	for i < len(height) {
		if len(stack) == 0 || height[i] > height[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var val int
			if len(stack) == 0 {
				val = height[temp] * i
			} else {
				val = height[temp] * (i - stack[len(stack)-1] - 1)
			}
			max = getMax(max, val)
		}
	}
	return max
}

func maximalRectangleStack(matrix [][]byte) int {
	max := 0
	height := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				height[j] += 1
			}
		}
		max = getMax(max, largestRectangleAreaStack(height))
	}
	return max
}

func getHeight(matrix [][]byte) [][]int {
	count := make([][]int, len(matrix))
	ok := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		count[i] = make([]int, len(matrix[0]))
		ok[i] = make([]bool, len(matrix[0]))
	}

	for i := len(matrix) - 1; i >= 0; i-- {
		for j := 0; j < len(matrix[0]); j++ {
			for k := i; k < len(matrix); k++ {
				if matrix[i][j] == '1' {
					if ok[k][j] == false {
						count[k][j] += 1
					}
				} else {
					ok[k][j] = true
				}
			}
		}
	}
	return count
}

func main() {
	input := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '1', '1', '0'},
	}

	fmt.Println(maximalRectangleStack(input))
}
