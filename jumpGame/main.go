package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	return step(nums, 0)
}

func step(nums []int, result int) int {
	if len(nums) <= 1 {
		return result
	}
	min := math.MaxUint32
	for i := 1; i <= nums[0] && i <= len(nums)-1; i++ {
		if temp := step(nums[i:], result+1); temp < min {
			min = temp
		}
	}
	return min
}

func jump2(nums []int) int {
	step := 0
	i := 0
	base := 0
	for i < len(nums)-1 {
		max := 0
		for j := base; j <= i; j++ {
			if nums[j]-(i-j) > max {
				max = nums[j] - (i - j)
				base = j
			}
		}
		i += max
		step++
	}
	return step
}

func main() {
	nums := []int{2, 1, 1, 3, 1, 1, 4}
	fmt.Println(jump2(nums))
}
