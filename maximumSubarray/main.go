package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	max := 0
	tm := 0
	temp := math.MinInt32
	for i := 0; i < len(nums); i++ {
		tm += nums[i]
		if tm < 0 {
			tm = 0
		}
		if tm > max {
			max = tm
		}
		if nums[i] > temp {
			temp = nums[i]
		}
	}
	if temp < 0 {
		return temp
	}
	return max
}

func getSum(nums []int, left int, right int) int {
	sum := 0
	for left <= right {
		sum += nums[left]
		left++
	}
	return sum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
