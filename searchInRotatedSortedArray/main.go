package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
	for nums[left] > nums[right] {
		if nums[left] > target && nums[right] < target {
			return -1
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		left++
		right--
	}
	return midSearch(nums, target, left, right)
}

func midSearch(nums []int, target int, left int, right int) int {
	pos := (right + left) / 2
	for nums[pos] != target && left != pos && right != pos {
		if nums[pos] < target {
			left = pos
		} else {
			right = pos
		}
		pos = (right + left) / 2
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	if nums[pos] != target {
		return -1
	}
	return pos
}

func main() {
	nums := []int{1, 3}
	target := 3
	fmt.Println(search(nums, target))
}
