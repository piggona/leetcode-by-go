package main

import "fmt"

func removeDuplicates(nums []int) int {
	left := 0
	pos := 1
	place := 1
	if len(nums) == 1 || len(nums) == 0 {
		return len(nums)
	}
	for pos < len(nums) {
		if nums[left] == nums[pos] {
			pos++
		} else {
			nums[place] = nums[pos]
			place++
			left++
			pos++
		}
	}

	return len(nums[:place])
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
