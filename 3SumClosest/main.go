package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	result := 65535
	sort.Ints(nums)
	length := 0
	for i := 0; i < len(nums)-2; i++ {
		middle := target - nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			length = nums[left] + nums[right]
			if length > middle {
				right--
			}
			if length < middle {
				left++
			}
			if abs(middle-length) < abs(result) {
				result = middle - length
			}
			if length == middle {
				return target
			}
		}
	}
	return target - result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	nums := []int{-1, 2, 1}
	fmt.Println(threeSumClosest(nums, 1))
}
