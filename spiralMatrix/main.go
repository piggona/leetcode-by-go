package main

import "fmt"

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

func main() {
	input := 5
	fmt.Println(generateMatrix(input))
}
