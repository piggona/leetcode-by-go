package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	if len(nums) < 4 {
		return result
	}
	sort.Ints(nums)
	resultMap := make(map[string]bool)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			left := j + 1
			right := len(nums) - 1
			m := target - nums[i] - nums[j]
			for left < right {
				if nums[left]+nums[right] < m {
					left++
				} else if nums[left]+nums[right] > m {
					right--
				} else {
					temp := []int{nums[i], nums[j], nums[left], nums[right]}
					tempB := []byte{byte(nums[i]), byte(nums[j]), byte(nums[left]), byte(nums[right])}
					if !resultMap[string(tempB)] {
						resultMap[string(tempB)] = true
						result = append(result, temp)
					}
					left++
					right--
				}
			}
		}
	}

	return result
}

func removeRep(slc []int) []int {
	result := []int{slc[0]}
	count := 0
	for i := 1; i < len(slc); i++ {
		if slc[i] != result[count] {
			result = append(result, slc[i])
			count++
		}
	}
	return result
}

func main() {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	target := 0
	fmt.Println(fourSum(nums, target))
}
