package main

import (
	"fmt"
	"math"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	pos := midSearch(nums, target)
	if pos == -1 {
		return []int{-1, -1}
	}
	left := pos
	right := pos
	for left >= 0 {
		if nums[left] == nums[pos] {
			left--
		} else {
			break
		}
	}
	for right < len(nums) {
		if nums[right] == nums[pos] {
			right++
		} else {
			break
		}
	}
	return []int{left + 1, right - 1}

}

func midSearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	for {
		mid = int(math.Floor(float64((left + right) / 2)))
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			break
		}
		if left > right {
			mid = -1
			break
		}
	}
	return mid
}

func main() {
	nums := []int{}
	fmt.Println(searchRange(nums, 8))
}
