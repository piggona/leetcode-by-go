package main

import (
	"fmt"
	"sort"
)

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}
	count := 1
	for i := len(nums); i > 0; i-- {
		count *= i
	}
	result := make([][]int, count)
	for i := 0; i < count; i++ {
		nextPermute(nums)
		temp := make([]int, len(nums))
		copy(temp, nums)
		result[i] = temp
	}
	return result
}

func nextPermute(nums []int) {
	pos := len(nums) - 2
	for pos-1 >= 0 && nums[pos] > nums[pos+1] {
		pos--
	}
	for i := len(nums) - 1; i > pos; i-- {
		if nums[pos] < nums[i] {
			nums[pos], nums[i] = nums[i], nums[pos]
			sort.Ints(nums[pos+1:])
			return
		}
	}
	sort.Ints(nums)

}

func main() {
	nums := []int{1}
	fmt.Println(permute(nums))
}
