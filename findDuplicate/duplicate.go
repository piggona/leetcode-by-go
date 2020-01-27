package findDuplicate

import "sort"

func findDuplicateBitmap(nums []int) int {
	bitmap := make([]bool, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		if bitmap[temp] {
			return temp
		}
		bitmap[temp] = true
	}
	return -1
}

func findDuplicateSort(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return 0
}

func findDuplicate(nums []int) int {
	left := 1
	right := len(nums) - 1
	for left <= right {
		mid := left + ((right - left) >> 1)
		cnt := 0
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}
