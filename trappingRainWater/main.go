package main

import "fmt"

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	result := 0
	dec := 0
	left := 0
	for i := 1; i < len(height); i++ {
		if height[i] >= height[left] {
			result += (i-left-1)*height[left] - dec
			left = i
			dec = 0
		} else {
			dec += height[i]
		}
	}
	last := height[left:]
	reverse(last)
	return result + trap(last)
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
	height := []int{4, 2, 3}
	fmt.Println(trap(height))
}
