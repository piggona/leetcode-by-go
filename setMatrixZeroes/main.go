package main

import "fmt"

func setZeroes(matrix [][]int) {
	rowMap := make(map[int]bool)
	colMap := make(map[int]bool)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				if _, exists := rowMap[row]; !exists {
					rowMap[row] = true
				}
				if _, exists := colMap[col]; !exists {
					colMap[col] = true
				}
			}
		}
	}
	for row, _ := range rowMap {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[row][i] = 0
		}
	}
	for col, _ := range colMap {
		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}
	}
}

func main() {
	input := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(input)
	fmt.Println(input)
}
