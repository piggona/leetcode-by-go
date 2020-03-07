package main

import (
	"fmt"
	"math"
)

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	line := 0
	col := 0
	candidate := 1

	pace := n - 1
	direction := 1
	recent := n - 1

	target := n * n

	for candidate <= target {
		result[line][col] = candidate
		candidate++
		switch direction {
		case 1:
			col++
		case 2:
			line++
		case 3:
			col--
		case 4:
			line--
		}
		recent--
		if recent == 0 {
			if direction == 4 {
				line++
				col++
				direction = 0
				pace -= 2
			}
			direction++
			recent = pace
		}
	}
	return result
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	direction := 0
	lmax := len(matrix)
	cmax := len(matrix[0])
	all := (lmax) * (cmax)
	result := []int{}
	count := 0
	l := 0
	c := 0
	for count < all-1 {
		switch direction {
		case 0:
			if c+1 < cmax && matrix[l][c+1] != math.MinInt32 {
				result = append(result, matrix[l][c])
				matrix[l][c] = math.MinInt32
				c++
				count++
			} else {
				direction = changeDirection(direction)
			}
		case 1:
			if l+1 < lmax && matrix[l+1][c] != math.MinInt32 {
				result = append(result, matrix[l][c])
				matrix[l][c] = math.MinInt32
				l++
				count++
			} else {
				direction = changeDirection(direction)
			}
		case 2:
			if c-1 >= 0 && matrix[l][c-1] != math.MinInt32 {
				result = append(result, matrix[l][c])
				matrix[l][c] = math.MinInt32
				c--
				count++
			} else {
				direction = changeDirection(direction)
			}
		case 3:
			if l-1 >= 0 && matrix[l-1][c] != math.MinInt32 {
				result = append(result, matrix[l][c])
				matrix[l][c] = math.MinInt32
				l--
				count++
			} else {
				direction = changeDirection(direction)
			}
		}
	}
	result = append(result, matrix[l][c])
	return result
}

func changeDirection(cur int) int {
	if cur < 3 {
		return cur + 1
	}
	return 0
}

func main() {
	// input := 5
	// fmt.Println(generateMatrix(input))
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(input))
}
