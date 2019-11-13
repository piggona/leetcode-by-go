package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := nums[0]
	count := 1
	left := 1
	pos := 1
	for pos < len(nums) {
		if nums[pos] != n {
			count = 1
			nums[left] = nums[pos]
			n = nums[pos]
			left++
			pos++
		} else if count == 2 {
			pos++
		} else {
			count++
			nums[left] = nums[pos]
			n = nums[pos]
			left++
			pos++
		}
	}
	return left
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
