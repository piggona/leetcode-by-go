package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	count := 1
	for i := len(nums); i > 0; i-- {
		count *= i
	}
	result := [][]int{}
	if len(nums) == 0 {
		return result
	}
	if len(nums) < 2 {
		return [][]int{nums}
	}

	dup := make(map[string]bool)
	for i := 0; i < count; i++ {
		nextPermute(nums)
		temp := make([]int, len(nums))
		copy(temp, nums)
		str := fmt.Sprint(temp)
		if !dup[str] {
			result = append(result, temp)
			dup[str] = true
		}
	}
	return result
}

func nextPermute(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := len(nums) - 1; j >= i+1; j-- {
				if nums[i] < nums[j] {
					nums[i], nums[j] = nums[j], nums[i]
					sort.Ints(nums[i+1:])
					return
				}
			}
		}
	}
	sort.Ints(nums)
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
