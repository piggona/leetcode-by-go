package main

import "sort"

func threeSumB(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var (
		result = [][]int{}
		cur    = 0
		left   = 1
		right  = len(nums) - 1
	)
	for cur < len(nums)-2 && nums[cur] <= 0 {
		if cur > 0 && nums[cur] == nums[cur-1] {
			continue
		}
		tcur := nums[cur]
		need := 0 - tcur
		for left < right {
			tl := nums[left]
			tr := nums[right]
			if tl+tr == need {
				result = append(result, append([]int{}, tcur, tl, tr))
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if tl+tr < need {
				left++
			} else {
				right--
			}
		}
		cur++
		left = cur + 1
		right = len(nums) - 1
	}
	return result
}
