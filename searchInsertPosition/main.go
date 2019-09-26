package main

import (
	"fmt"
	"math"
)

func searchInsert(nums []int, target int) int {
	return midSearch(nums, target)
}

func midSearch(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	mid := 0
	for {
		mid = int(math.Floor(float64(left+right) / 2))
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			break
		}
		if left > right {
			mid = left
			break
		}
	}
	return mid
}

func main() {
	nums := []int{}
	target := 7
	fmt.Println(searchInsert(nums, target))
}
