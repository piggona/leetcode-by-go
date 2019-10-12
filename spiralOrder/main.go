package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	col := len(matrix) - 1
	line := len(matrix[0]) - 1
	direction := false // true为平行，false为垂直
	spread := true     // true为正方向，false为反方向
	i := 0
	j := col
	result := []int{}
	result = append(result, matrix[0]...)

	for col != 0 && line != 0 {
		if direction {
			if spread {
				if j+line+1 > len(matrix[0])-1 {
					result = append(result, matrix[i][j+1:]...)
				} else {
					result = append(result, matrix[i][j+1:j+line+1]...)
				}
				j = j + line
			} else {
				result = append(result, matrix[i][j-line:j]...)
				j = j - line
			}
			direction = !direction
			spread = !spread
		} else {
			if spread {
				if i+col+1 > len(matrix)-1 {
					result = append(result, matrix[i+1:][j]...)
				} else {
					result = append(result, matrix[i+1 : i+col+1][j]...)
				}
				i = i + col
			} else {
				result = append(result, matrix[i-col : i][j]...)
				i = i - col
			}
			direction = !direction
			spread = !spread
		}
	}
	return result
}

func main() {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(input))
}
