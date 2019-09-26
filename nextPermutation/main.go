package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			pos := len(nums) - 1
			for nums[pos] <= nums[i] {
				pos--
			}
			nums[i], nums[pos] = nums[pos], nums[i]
			sort.Ints(nums[i+1:])
			return
		}
	}
	sort.Ints(nums)
}

func main() {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
