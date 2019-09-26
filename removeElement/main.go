package main

import "fmt"

func removeElement(nums []int, val int) int {
	place := -1
	pos := 0
	for pos < len(nums) && nums[pos] != val {
		pos++
	}
	if pos >= len(nums) {
		return len(nums)
	}
	place = pos
	pos++
	for pos < len(nums) {
		if nums[pos] != val {
			nums[place] = nums[pos]
			place++
		}
		pos++
	}
	return len(nums[:place])
}

func main() {
	nums := []int{8}
	fmt.Println(removeElement(nums, 8))
}
