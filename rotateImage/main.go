package main

import "fmt"

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		reverse(matrix[i])
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)-i; j++ {
			matrix[i][j], matrix[len(matrix)-j-1][len(matrix)-i-1] = matrix[len(matrix)-j-1][len(matrix)-i-1], matrix[i][j]
		}
	}
}

func reverse(nums []int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
