package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := 0
	area := 0
	for left < right {
		smaller := min(height[left], height[right])
		area = smaller * (right - left)
		if area > result {
			result = area
		}
		if height[left] < height[right] {
			ruler := height[left]
			left++
			for ruler > height[left] {
				left++
			}
		} else {
			ruler := height[right]
			right--
			for ruler > height[right] {
				right--
			}
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	s := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(s))
}
