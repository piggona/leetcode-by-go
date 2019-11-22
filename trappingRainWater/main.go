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

func trapStack(height []int) int {
	stack := []int{}
	result := 0
	i := 0
	for i < len(height) {
		if len(stack) == 0 || height[i] <= height[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				continue
			}
			left := stack[len(stack)-1]
			high := 0
			if height[i] > height[left] {
				high = height[left]
			} else {
				high = height[i]
			}
			result += (high - height[temp]) * (i - left - 1)
		}
	}
	return result
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trapStack(height))
}
